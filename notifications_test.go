package go_wypok

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// As notifications are nearly impossible to be tested automatically, you need a help everytime you run tests.
func TestWykopHandler_TestNotification(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	wh.LoginToWypok()

	//success, err := wh.readNotification(9299079331)
	//
	//fmt.Println("ReadNotification success: ", success)
	//fmt.Println("ReadNotification err: ", err)

	//notifications, err := wh.getNotifications(1)
	//fmt.Println("notifications: ", notifications)
	//fmt.Println("notificationsErr: ", err)
	//
	//notifications, err = wh.getHashTagsNotifications(1)
	//fmt.Println("#notifications: ", notifications)
	//fmt.Println("#notificationsErr: ", err)
	//
	//counter, err := wh.getNotificationsCount()
	//fmt.Println("NotificationsCounter: ", counter)
	//fmt.Println("NotificationsErr: ", err)
	//
	//counter, err = wh.getHashTagsNotificationsCount()
	//fmt.Println("#NotificationsCounter: ", counter)
	//fmt.Println("#NotificationsErr: ", err)
}

func TestWykopHandler_getNotificationsUrls(t *testing.T) {
	wh := WykopHandler{}
	wh.authResponse.Userkey = "123"
	wh.appKey = "abc"

	notificationsUrl := getNotificationsUrl(1, &wh)
	notificationsCountUrl := getNotificationsCountUrl(&wh)
	hashTagsNotificationsUrl := getHashTagsNotificationsUrl(1, &wh)
	hashTagsNotificationsCountUrl := getHashTagsNotificationsCountUrl(&wh)
	readNotificationsUrl := getReadNotificationsUrl(&wh)
	readHashTagsNotificationsUrl := getReadHashTagsNotificationsUrl(&wh)
	markNotificationAsReadUrl := getMarkNotificationAsReadUrl(1, &wh)

	assert.Equal(t, notificationsUrl, "https://a.wykop.pl/MYWYKOP/Notifications/appkey/abc/userkey/123/page/1")
	assert.Equal(t, notificationsCountUrl, "https://a.wykop.pl/MYWYKOP/NotificationsCount/appkey/abc/userkey/123")
	assert.Equal(t, hashTagsNotificationsUrl, "https://a.wykop.pl/MYWYKOP/HashTagsNotifications/appkey/abc/userkey/123/page/1")
	assert.Equal(t, hashTagsNotificationsCountUrl, "https://a.wykop.pl/MYWYKOP/HashTagsNotificationsCount/appkey/abc/userkey/123")
	assert.Equal(t, readNotificationsUrl, "https://a.wykop.pl/MYWYKOP/ReadNotifications/appkey/abc/userkey/123")
	assert.Equal(t, readHashTagsNotificationsUrl, "https://a.wykop.pl/MYWYKOP/ReadHashTagsNotifications/appkey/abc/userkey/123")
	assert.Equal(t, markNotificationAsReadUrl, "https://a.wykop.pl/MYWYKOP/MarkAsReadNotification/1/appkey/abc/userkey/123")
}
