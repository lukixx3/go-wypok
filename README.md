[![pipeline status](https://gitlab.com/agilob/go-wypok/badges/master/pipeline.svg)](https://gitlab.com/agilob/go-wypok/commits/master)
[![coverage report](https://gitlab.com/agilob/go-wypok/badges/master/coverage.svg)](https://gitlab.com/agilob/go-wypok/commits/master)

# go-wypok
Obsługa API wykopu w go

```
go get gitlab.com/agilob/go-wypok
```

# Wspierane metody API w kolejnosi identycznej jak na stronie [dokumentacji](https://www.wykop.pl/dla-programistow/dokumentacja/)

W nawiasie podana jest metoda obsługująca dane endpoint.
Jeśli endpoint ma opcjonalne pola (np. embed), istnieje druga metoda o podobnej nazwie (`PostEntry` i `PostEntryWithEmbeddedContent`)

1. Comments
 - [x] Add (AddComment - for comments without embedded item, pass empty string ("") as embedUrl param, also for "main" comment, so a  comment which starts a discussion, pass 0 as parentId)
 - [x] Plus (PlusComment)
 - [x] Minus (MinusComment)
 - [x] Edit (EditComment)
 - [x] Delete (DeleteComment)
2. Link
 - [ ] Index
 - [ ] Dig
 - [ ] Cancel
 - [ ] Bury
 - [ ] Comments
 - [ ] Reports
 - [ ] Digs
 - [ ] Related
 - [ ] Buryreasons
 - [ ] Observe
 - [ ] Favorite
3. Links
 - [x] Promoted (GetMainPageLinks)
 - [x] Upcoming (GetUpcomingLinks)
4. Popular
 - [ ] Promoted
 - [ ] Upcoming
5. Profile
 - [x] Index (GetProfile)
 - [x] Added (GetProfileAdded)
 - [?] Groups (undocumented endpoint, returns nothing)
 - [x] Published (GetProfilePublished)
 - [x] Commented (GetProfileCommented)
 - [x] Comments (GetProfileComments)
 - [x] Digged (GetProfileDigged)
 - [x] Buried (GetProfileBuried)
 - [x] Observe (ObserveProfile)
 - [x] Unobserve (UnobserveProfile)
 - [x] Block (BlockProfile)
 - [x] Unblock (UnblockProfile)
 - [x] Followers (ProfileFollowers)
 - [x] Followed (ProfileFollowed)
 - [x] Favorites (GetProfileFavorites)
 - [x] Entries (GetProfileEntries)
 - [x] EntriesComments (GetProfileEntriesComments)
 - [ ] Related
6. Search
 - [ ] Index
 - [ ] Links
 - [ ] Entries
 - [ ] Profiles
7. User
 - [x] Login (LoginToWypok)
 - [ ] Favorites
 - [ ] Observed
 - [ ] Tags
 - [ ] Connect
8. Top
 - [ ] Index
 - [ ] Date
 - [ ] Hits
9. Add
 - [ ] Index
10. Related
 - [ ] Plus
 - [ ] Minus
 - [ ] Add
11. Mywykop
 - [ ] Index
 - [ ] Tags
 - [ ] Users
 - [ ] Observing // deprecated
 - [ ] Mine // deprecated
 - [ ] Received // deprecated
 - [x] Notifications (getNotifications)
 - [x] NotificationsCount (getNotificationsCount)
 - [x] HashTagsNotifications (getHashTagsNotifications)
 - [x] HashTagsNotificationsCount (getHashTagsNotificationsCount)
 - [x] ReadNotifications (readNotifications)
 - [x] ReadHashTagsNotifications (readHashTagsNotifications)
 - [x] MarkAsReadNotification (readNotification)
12. Entries
 - [x] Index (GetEntry)
 - [x] Add (PostEntry)
 - [x] Edit (EditEntry)
 - [x] Delete (DeleteEntry)
 - [x] AddComment (AddEntryComment)
 - [x] EditComment (EditEntryComment)
 - [x] DeleteComment (DeleteEntryComment)
 - [x] Vote (UpvoteEntry)
 - [x] Unvote (UnvoteEntry)
 - [x] Favorite (FavoriteEntry)
 - [ ] Polls (available only in APiv2)
13. Rank
 - [x] Index
14. Observatory
 - [ ] Votes
 - [ ] Comments
 - [ ] Entries
 - [ ] EntriesComments
15. Favorites
 - [x] Index (GetFavoritesListLinks)
 - [x] Comments (GetFavoritesComments)
 - [x] Entries (GetFavoritesEntries)
 - [x] Lists (GetFavoritesLists)
16. Stream
 - [x] Index (GetStreamEntries)
 - [x] Hot
   - GetStreamLast6HoursHotEntries
   - GetStreamLast12HoursHotEntries
   - GetStreamLast24HoursHotEntries
17. Tag
 - [ ] Index
 - [ ] Links
 - [x] Entries (GetEntriesFromTag)
 - [ ] Observe
 - [ ] Unobserve
 - [ ] Block
 - [ ] Unblock
18. PM
 - [x] ConversationsList (GetConversationsList)
 - [x] Conversation (GetConversation)
 - [x] SendMessage (SendPrivateMessageTo, SendPrivateMessageWithEmbeddedUrlTo)
 - [x] DeleteConversation (DeleteConversation)
19. Tags
 - [ ] Index
