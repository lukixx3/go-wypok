[![pipeline status](https://gitlab.com/agilob/go-wypok/badges/master/pipeline.svg)](https://gitlab.com/agilob/go-wypok/commits/master)
[![coverage report](https://gitlab.com/agilob/go-wypok/badges/master/coverage.svg)](https://gitlab.com/agilob/go-wypok/commits/master)

# go-wypok
Obsługa API wykopu w go

# Wspierane metody API w kolejnosi identycznej jak na stronie [dokumentacji](https://www.wykop.pl/dla-programistow/dokumentacja/)

W nawiasie podana jest metoda obsługująca dane endpoint.
Jeśli endpoint ma opcjonalne pola (np. embed), istnieje druga metoda o podobnej nazwie (`PostEntry` i `PostEntryWithEmbeddedContent`)

1. Comments
 - [ ] Add
 - [ ] Plus
 - [ ] Minus
 - [ ] Edit
 - [ ] Delete
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
 - [ ] Added
 - [ ] Groups
 - [ ] Published
 - [ ] Commented
 - [x] Comments (GetProfileComments)
 - [ ] Digged
 - [ ] Buried
 - [ ] Observe
 - [ ] Unobserve
 - [ ] Block
 - [ ] Unblock
 - [ ] Followers
 - [ ] Followed
 - [ ] Favorites
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
 - [ ] Observing
 - [ ] Mine
 - [ ] Received
 - [ ] Notifications
 - [ ] NotificationsCount
 - [ ] HashTagsNotifications
 - [ ] HashTagsNotificationsCount
 - [ ] ReadNotifications
 - [ ] ReadHashTagsNotifications
 - [ ] MarkAsReadNotification
12. Entries
 - [x] Index (GetEntry)
 - [x] Add (PostEntry)
 - [x] Edit (EditEntry)
 - [x] Delete (DeleteEntry)
 - [x] AddComment (AddEntryComment)
 - [ ] EditComment
 - [x] DeleteComment (DeleteEntryComment)
 - [x] Vote (UpvoteEntry)
 - [ ] Unvote
 - [ ] Favorite
13. Rank
 - [x] Index
14. Observatory
 - [ ] Votes
 - [ ] Comments
 - [ ] Entries
 - [ ] EntriesComments
15. Favorites
 - [ ] Index
 - [ ] Comments
 - [ ] Entries
 - [ ] Lists
16. Stream
 - [ ] Index
 - [ ] Hot
17. Tag
 - [ ] Index
 - [ ] Links
 - [x] Entries (GetEntriesFromTag)
 - [ ] Observe
 - [ ] Unobserve
 - [ ] Block
 - [ ] Unblock
18. PM // In progress
 - [ ] ConversationsList
 - [ ] Conversation
 - [ ] SendMessage
 - [ ] DeleteConversation
19. Tags
 - [ ] Index
