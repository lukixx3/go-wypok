package go_wypok

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

const (
	BASE_NOTIFICATIONS_URL = "https://a.wykop.pl/MYWYKOP/%s/appkey/%s/userkey/%s"
)

type Notification struct {
	Id           uint          `json:"id"`
	Author       string        `json:"author"`
	Avatar       string        `json:"author_avatar"`
	AvatarMedium string        `json:"author_avatar_med"`
	AvatarLow    string        `json:"author_avatar_lo"`
	AvatarBig    string        `json:"author_avatar_big"`
	AuthorGroup  string        `json:"author_group"`
	AuthorSex    string        `json:"author_sex"`
	Date         WypokShitDate `json:"date"`
	Content      string        `json:"body"`
	Type         string        `json:"type"`
	Url          string        `json:"url"`
}

func (wh *WykopHandler) getNotifications(page uint) (notifications []Notification, wykopError *WykopError) {
	urlAddress := getNotificationsUrl(page, wh)

	_, responseBody, _ := gorequest.New().Post(urlAddress).
		Set(contentType, mediaTypeFormType).
		Set(apiSignHeader, wh.hashRequest(urlAddress)).
		End()

	wykopError = wh.getObjectFromJson(responseBody, &notifications)

	return
}

func (wh *WykopHandler) getNotificationsCount() (counter uint, wykopError *WykopError) {
	urlAddress := getNotificationsCountUrl(wh)

	_, responseBody, _ := gorequest.New().Post(urlAddress).
		Set(contentType, mediaTypeFormType).
		Set(apiSignHeader, wh.hashRequest(urlAddress)).
		End()

	wykopError = wh.getObjectFromJson(responseBody, &counter)

	return
}

func (wh *WykopHandler) getHashTagsNotifications(page uint) (notifications []Notification, wykopError *WykopError) {
	urlAddress := getHashTagsNotificationsUrl(page, wh)

	_, responseBody, _ := gorequest.New().Post(urlAddress).
		Set(contentType, mediaTypeFormType).
		Set(apiSignHeader, wh.hashRequest(urlAddress)).
		End()

	wykopError = wh.getObjectFromJson(responseBody, &notifications)

	return
}

func (wh *WykopHandler) getHashTagsNotificationsCount() (counter uint, wykopError *WykopError) {
	urlAddress := getHashTagsNotificationsCountUrl(wh)

	_, responseBody, _ := gorequest.New().Post(urlAddress).
		Set(contentType, mediaTypeFormType).
		Set(apiSignHeader, wh.hashRequest(urlAddress)).
		End()

	wykopError = wh.getObjectFromJson(responseBody, &counter)

	return
}

func (wh *WykopHandler) readNotifications() (wykopError *WykopError) {
	urlAddress := getReadNotificationsUrl(wh)

	_, responseBody, _ := gorequest.New().Post(urlAddress).
		Set(contentType, mediaTypeFormType).
		Set(apiSignHeader, wh.hashRequest(urlAddress)).
		End()

	wykopError = wh.getObjectFromJson(responseBody, nil)

	return
}

func (wh *WykopHandler) readHashTagsNotifications() (wykopError *WykopError) {
	urlAddress := getReadHashTagsNotificationsUrl(wh)

	_, responseBody, _ := gorequest.New().Post(urlAddress).
		Set(contentType, mediaTypeFormType).
		Set(apiSignHeader, wh.hashRequest(urlAddress)).
		End()

	wykopError = wh.getObjectFromJson(responseBody, nil)

	return
}

func (wh *WykopHandler) readNotification(notificationId uint) (result bool, wykopError *WykopError) {
	urlAddress := getMarkNotificationAsReadUrl(notificationId, wh)

	_, responseBody, _ := gorequest.New().Post(urlAddress).
		Set(contentType, mediaTypeFormType).
		Set(apiSignHeader, wh.hashRequest(urlAddress)).
		End()
	result = false
	wykopError = wh.getObjectFromJson(responseBody, &result)

	return
}

func getNotificationsUrl(page uint, wh *WykopHandler) string {
	return fmt.Sprintf(BASE_NOTIFICATIONS_URL, "Notifications", wh.appKey, wh.authResponse.Userkey) + fmt.Sprintf("/page/%d", page)
}

func getNotificationsCountUrl(wh *WykopHandler) string {
	return fmt.Sprintf(BASE_NOTIFICATIONS_URL, "NotificationsCount", wh.appKey, wh.authResponse.Userkey)
}

func getHashTagsNotificationsUrl(page uint, wh *WykopHandler) string {
	return fmt.Sprintf(BASE_NOTIFICATIONS_URL, "HashTagsNotifications", wh.appKey, wh.authResponse.Userkey) + fmt.Sprintf("/page/%d", page)
}

func getHashTagsNotificationsCountUrl(wh *WykopHandler) string {
	return fmt.Sprintf(BASE_NOTIFICATIONS_URL, "HashTagsNotificationsCount", wh.appKey, wh.authResponse.Userkey)
}

func getReadNotificationsUrl(wh *WykopHandler) string {
	return fmt.Sprintf(BASE_NOTIFICATIONS_URL, "ReadNotifications", wh.appKey, wh.authResponse.Userkey)
}

func getReadHashTagsNotificationsUrl(wh *WykopHandler) string {
	return fmt.Sprintf(BASE_NOTIFICATIONS_URL, "ReadHashTagsNotifications", wh.appKey, wh.authResponse.Userkey)
}

func getMarkNotificationAsReadUrl(notificationId uint, wh *WykopHandler) string {
	return fmt.Sprintf(BASE_NOTIFICATIONS_URL, fmt.Sprintf("MarkAsReadNotification/%d", notificationId), wh.appKey, wh.authResponse.Userkey)
}
