package usecases

import (
	"avito-flats/internal/adapters/output/repositories/subscribe"
	"avito-flats/internal/adapters/output/services"
	"avito-flats/internal/domain/entities"
	"context"
	"log"
	"strconv"
	"time"
)

type SubscribeUsecase struct {
	subscribe      subscribe.Repository
	houseid        entities.HouseID
	notificationCh chan notification
	sender         *services.Sender
}

type notification struct {
	HouseID int
	Email   string
}

func NewSubscribeUsecase(subscribe subscribe.Repository, houseid entities.HouseID) *SubscribeUsecase {
	u := &SubscribeUsecase{
		subscribe:      subscribe,
		houseid:        houseid,
		notificationCh: make(chan notification, 100), // буферизованный канал
		sender:         services.NewSender(),
	}
	go u.processNotifications()
	return u
}

func (u *SubscribeUsecase) SubscribeToHouseUpdates(ctx context.Context, houseID entities.HouseID, email string) error {
	// Проверка существования дома
	_, err := u.subscribe.GetSubscribers(ctx, houseID)
	if err != nil {
		return err
	}

	// Добавление подписки
	return u.subscribe.SubscribeOnHouse(ctx, houseID, email)
}

func (u *SubscribeUsecase) NotifySubscribers(ctx context.Context, houseID entities.HouseID) error {
	subscribers, err := u.subscribe.GetSubscribers(ctx, houseID)
	if err != nil {
		return err
	}

	for _, email := range subscribers {
		u.notificationCh <- notification{HouseID: int(houseID), Email: email}
	}

	return nil
}

func (u *SubscribeUsecase) processNotifications() {
	for n := range u.notificationCh {
		// Создаем новый контекст с таймаутом для каждой отправки
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		err := u.sender.SendEmail(ctx, n.Email, "В доме с ID "+strconv.Itoa(n.HouseID)+" появились новые квартиры!")
		if err != nil {
			// Обработка ошибки отправки (например, логирование)
			log.Printf("Ошибка отправки уведомления: %v", err)
		}
		cancel() // Отменяем контекст после использования
	}
}
