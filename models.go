package go_wypok

import (
	"strconv"
	"time"
)

type Profile struct {
	ID              int64
	Login           string
	Email           string
	PublicEmail     string `json:"public_email"`
	Name            string
	Www             string
	Jabber          string
	Gg              int
	City            string
	About           string
	AuthorGroup     int `json:"author_group"`
	LinksAdded      int `json:"links_added"`
	LinksPublished  int `json:"links_published"`
	Comments        int
	Rank            int
	Followers       int
	Following       int
	Entries         int
	EntriesComments int `json:"entries_comments"`
	Diggs           int
	Buries          int
	RelatedLinks    int `json:"related_links"`
	Groups          int
	Sex             string
	Avatar          string
	AvatarLo        string `json:"avatar_lo"`
	AvatarMed       string `json:"avatar_med"`
	AvatarBig       string `json:"avatar_big"`
	IsObserved      bool   `json:"is_observed"`
}

type LinkComment struct {
	ID              int    `json:"id"`
	Date            string `json:"date"`
	Author          string `json:"author"`
	AuthorGroup     int    `json:"author_group"`
	AuthorAvatar    string `json:"author_avatar"`
	AuthorAvatarBig string `json:"author_avatar_big"`
	AuthorAvatarMed string `json:"author_avatar_med"`
	AuthorAvatarLo  string `json:"author_avatar_lo"`
	AuthorSex       string `json:"author_sex"`
	VoteCount       int    `json:"vote_count"`
	VoteCountPlus   int    `json:"vote_count_plus"`
	VoteCountMinus  int    `json:"vote_count_minus"`
	Body            string `json:"body"`
	Source          string `json:"source"`
	ParentID        int    `json:"parent_id"`
	Status          string `json:"status"`
	CanVote         bool   `json:"can_vote"`
	UserVote        bool   `json:"user_vote"`
	Blocked         bool   `json:"blocked"`
	Deleted         bool   `json:"deleted"`
	Embed           Embed  `json:"embed"`
	Type            string `json:"type"`
	App             string `json:"app"`
	UserFavorite    bool   `json:"user_favorite"`
	ViolationURL    string `json:"violation_url"`
	Link            Link
}

type WykopShitUserVote string

// when user is NOT logged in, user_vote field will not be provided
// when user is logged in, user_vote might have value of "bury" or "dig"
// if user made such action in the past, however, when user didn't vote
// on the link wypok.pl provides user_vote as false (see no quotes), which
// golang unmarshaller tries to read as `bool` not `string` and panics.
// This wrapper type ensures that bool will be converted to string and string treated as string.
func (value *WykopShitUserVote) UnmarshalJSON(data []byte) error {
	asString := string(data)
	if asString == "dig" {
		*value = "dig"
	} else if asString == "bury" {
		*value = "bury"
	} else {
		*value = "false"
	}
	return nil
}

type WypokShitDate struct {
	time.Time
}

func (self *WypokShitDate) UnmarshalJSON(b []byte) (err error) {
	s := string(b)

	// Get rid of the quotes "" around the value.
	// A second option would be to include them
	// in the date format string instead, like so below:
	//   time.Parse(`"`+time.RFC3339Nano+`"`, s)
	s = s[1 : len(s)-1]

	t, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		t, err = time.Parse("2006-01-02 15:04:05", s)
	}
	self.Time = t
	return
}

type Voter struct {
	Author          string
	AuthorAvatar    string        `json:"author_avatar"`
	AuthorAvatarBig string        `json:"author_big"`
	AuthorAvatarMed string        `json:"author_med"`
	AuthorAvatarLo  string        `json:"author_lo"`
	AuthorSex       string        `json:"author_sex"`
	AuthorGroup     int           `json:"author_group"`
	Date            WypokShitDate `json:"date"`
}

type Embed struct {
	EmbedType string `json:"type"`
	Preview   string
	Url       string
	Source    string
	Plus18    bool `json:"plus_18"`
}

type TagsEntries struct {
	Meta  Meta
	Items []Entry
}

type Meta struct {
	Tags       string
	IsObserved bool `json:"is_observed"`
	IsBlocked  bool `json:"is_blocked"`
	Counters   Counters
}

type Counters struct {
	Total   int
	Entries int
	Links   int
}

type VoteResponse struct {
	Vote int
	//Voters string
}

type FavoriteResponse struct {
	UserFavorite bool `json:"user_favorite"`
}

type AuthenticationResponse struct {
	Login        string
	Email        string
	ViolationUrl string `json:"violation_url"`
	Userkey      string
}

type EntryResponse struct {
	Id int
}

type CommentResponse struct {
	Id int
}

// WTF wykop.pl API is so retarded. Everywhere it returns int as Id on any response class
// by comment/entry submitted response returns id in quotes "", literally - integer inside quotes,
// which makes it a fucking string, so these struct wrapper classes unmarshall it to string
// then the struct is converted to a new one, where id is an int. bravo białov. bravo.
// thanks for reading this rant, please enjoy comment at the bottom of this file.
type fuckingStringEntryResponse struct {
	// for some stupid reason wypok returns string here
	Id string `json:"id"`
}

type fuckingStringCommentResponse struct {
	// and here, it should be int
	Id string `json:"id"`
}

func newEntryResponse(fuckingStringResponse fuckingStringEntryResponse) EntryResponse {
	entryResponse := EntryResponse{}
	theId, err := strconv.Atoi(fuckingStringResponse.Id)
	if err != nil {
		entryResponse.Id = 0
	} else {
		entryResponse.Id = theId
	}
	return entryResponse
}

func newCommentResponse(fuckingStringResponse fuckingStringCommentResponse) CommentResponse {
	commentResponse := CommentResponse{}
	theId, err := strconv.Atoi(fuckingStringResponse.Id)
	if err != nil {
		commentResponse.Id = 0
	} else {
		commentResponse.Id = theId
	}
	return commentResponse
}

type WykopError struct {
	ErrorObject WykopErrorMessage `json:"error"`
}

type WykopErrorMessage struct {
	Code    int
	Message string
}

// Michał Białek kończył nocną wartę w serwerowni wykopu. Za oknem zadłużonej willi poznańskie koziołki ocierały się
// częściami, których Białek wolałby nigdy nie mieć. Przypomniał sobie o żonie, którą widywał głównie w niedzelne poranki.
// Stękanie ćwiczącego lewicę Kinera dobiegało zza rzędu wykopowych monitorów. Białek automatycznie podłożył ociekające
// potem dźwięki pod obraz żony osadzonej na knadze Kinera. Intuicja podpowiadała mu, że lekko otyły kolega z pracy
// przebiera palcami po najkrótszej części ciała oglądając zdjęcia białkowej połowicy na fejsbuku. Nie mylił się. Czuł jednak dziwaczną dumę połączoną z rozbawieniem, które przyniosła mu owa wizja.
// ,g,Puk puk".
// Pierwsza myśl - Elfik32. Kurwiszcze, które zrobiłoby wszystko za status moderatora w serwisie.
// Białek poczuł się ważny. Myśl o zdradzie żony przerodziła się w pewność, że w tę sobotnią noc powstaje dziecko,
// na które po badaniach DNA nie musiałby wyłożyć ani grosza. Zatliła się w jego biednym umyśle żądza zemsty.
// Ocknąwszy się wstał i nonszalancko otworzył drzwi.
// Elfik32. Czarna owca rodu Steckich stała w progu oparta o framugę ze swoim kurewskim uśmieszkiem woźnej.
// Baletki, krótka spódniczka i motzno zarysowany dekolt jaśniały kontrastującą z nocnym krajobrazem bielą.
// - wstawiona jak zwykle - pomyślał Białek.
// Elfik położyła palec na ustach. Białek domyślił się, że odseparowany od świata zewnętrznego Kiner nie usłyszał
// pukania aktywnej wykopowiczki, która właśnie przyniosła im kanapki. Spod granicy niemieckiej.
// Elfik sprawnie zzuła obuwie i figlarnie mrugając ruszyła w stronę Kinera. Białek podążał wzrokiem za opalonymi
// stópkami zmierzającymi w stronę McKinera, nie mogąc powstrzymać wewnętrznego rozbawienia. Za chwilę miał wyjść na jaw fakt,
// którego nikt z pracowników wykopu osobiście nie widział, choć był świadom jego istnienia. Fakt, który miał zburzyć
// spokój Kinera na zawsze.
// Elfik wczołgała się pod biurko i sprawnie wyskoczyła po drugiej stonie. Kiner wybałuszył oczy i odskoczył w tył
// z naprężonym kutasem w ręku.
// - Ale... Spojrzenia Białka i Kinera spotkały się. Kiner był zażenowany całą sytuacją, o czym świadczył wyraźny rumieniec
// na jego aryjskiej twarzy.
// Śmiech Elfika rozniósł się po pomieszczeniu.
// - Mała pała jak na administratora. - powiedziała Elfik, ledwie powstrzymując śmiech. Kiner spąsowiał bardziej.
// Ręce machinalnie powędrowały w stronę rozporka, gdy podchmielona Elfik rzuciła się w tę samą stronę.
// - Zostaw. - powiedziała stanowczo.
// Zaskoczony Kiner wypuścił pytonga z ręki. Dziesięciocentymetrowy organ bezwładnie opadł lekko kołysząc się na boki,
// gdy Elfik doskoczyła do sparaliżowanego grubaska niczym wygłodniała kura i stanowczym ruchem opuściła nieco za duże
// spodnie. Od czasu nieudanego eksperymentu z rurkami Kiner powrócił bowiem do starych nawyków, co ułatwiło wykopowiczce zadanie.
// - Jesteś pijana. - wystękał, gdy Elfik chwyciła go za lekko przywiędniętą pałę i chichocząc zaczęła prowadzić w
// stronę Białka, który obserwował sytuację z zaciekawieniem. Programista nie protestował jednak zbyt zaciekle.
// Był to pierwszy raz, kiedy kobieca ręka spoczęła na jego wypustce. Było to niewątpliwie ciekawe doznanie,
// gdyż sam Kiner zdawał się zapomnieć o dziwnych okolicznościach, w jakich doszło do tego przełomowego momentu.
// Białek nie mógł już powstrzymać rozbawienia, obserwując zesztywniałego Kinera kroczącego za dzierżącą jego
// orzeszek Elfikiem. Wybuchnął serdecznym śmiechem, mimowolnie puszczając krótkiego bąka o dosyć wysokiej tonacji.
// - Przepraszam, - rzekł Białek przez łzy - ale nie bardzo rozumiem sytuację.
// Elfik puściła kinerowe przyrodzenie i kołysząc biodrami powoli podeszła do administratora o podkrążonych od pracoholizmu
// oczach, zalotnie kręcąc loczek wydłubany spod natapirowanej burzy blond włosów.
// - Nie planowałam tego. Zawsze byłam spontaniczna. - wyszeptała prowokacyjnie, gmerając już teraz palcami w okolicy
// guzików różowej koszuli Białka.
// Kiner zastygł jak posąg w centrum serwerowni. Nie zdawał sobie sprawy, jak komicznie wyglądał z opuszczonymi do kostek
// spodniami kupionymi przez mamę w second-handzie, z włosami łonowymi w nieładzie i zaczerwienioną od uścisku knagą
// smętnie zwisającą między otłuszczonymi udami. Poczuł ukłucie zazdrości widząc pierwszą kobietę, której pozwolił się dotknąć,
// z zapałem liżącą opalony tors Białka. To on powinien być na jej miejscu.
// Białek zamknął oczy, czując wilgotne pociągnięcia elfikowego języka po swojej klacie. Nie była to zdrada, był to gwałt.
// Stąd też, domyślając się dalszego przebiegu sytuacji, nie miał żadnych wyrzutów sumienia. Aby jednak im zapobiec,
// wyobraził sobie że jego żona jest właśnie posuwana przez murzyna.
// ,,Co za kurwa" - pomyślał. Czuł się całkowicieoczyszczony z zarzutów.
// Elfik przeszła do lizania twarzy, by w końcu zbliżyć się do ucha Białka.
// - Wiesz, do czego tasował twój kolega? - szeptnęła, po kurewsku przenosząc wzrok na jego twarz.
// - Do zdjęć twoich przeróbek zrobionych przez...
// - Nie kończ. - przerwał jej Białek. Nie chciał by słowo codziennie odmieniane przez przypadki w miejscu
// jego pracy ostudziło podniecenie. Na moment jednak otrzeźwiał i odepchnął rozpaloną Stecką od siebie.
// - Czego tak właściwie chcesz, hm?
// - zapytał, badawczo spoglądając na Elfika. Nie doczekał jednak odpowiedzi,
// gdyż Kiner niespodziewanie zwinnie, biorąc pod uwagę jego warunki fizyczne, chwycił drukarkę Samsunga i ogłuszył
// Elfika celnym uderzeniem w tył głowy. Elfik bezwładnie osunęła się na ziemie, potwierdzając swoją renomę kobiety upadłej. W sekundę
// później Kiner stanął z roznegliżowanym Białkiem twarzą w twarz, oko w oko. Ich chuje równiez były całkiem blisko.
// Zszokowany Białek nie rozpoznawał nieśmiałego dotąd kolegi. Coś w nim zdecydowanie pękło, a w spojrzeniu
// programisty była niewidoczna dotąd determinacja.
// - Zerżnij mnie. - wycedził Kiner.
// - Zerżnij mnie motzno w odbyt. - powtórzył. Powieka nawet nie drgnęła podczas
// wypowiadania tych słów.
// Białek w swoim zaskoczeniu wydał niezidentyfikowany dźwięk, lecz Kiner natychmiast położył mu palec na ustach.
// - Nikt się nie dowie. Ta kurwa Elfik i tak ci już powiedziała. Nie mam nic do stracenia. - powiedział.
// Nie czekając na reakcję Białka, przeszedł od słów do czynów.
// Świat Białka wywrócił się do góry nogami, starał się jednak chłodno kalkulować, jak wiele nieoczekiwana przygoda
// gejowska mogła zmienić w imidżu samca alfa, na który tak ciężko pracował. A Kiner nie żartował.
// Był silniejszy od niego, co potwierdził brutalnie atakując Elfika kilkanaście sekund wcześniej.
// Setki podobnych myśli przelatywały mu przez głowę, gdy rudawy programista delikatnie rozchylił mu wargi,
// by włożyć mięciutki palec do ust i wymusić odruch ssania.
// Patrząc odważnie w oczy Białka, Kiner zrobił kilka kroków w tył i zdjął t-shirt jednym pewnym ruchem.
// Biała delikatna skóra i puchate sutki Kinera były dziwacznie atrakcyjne, choć jeszcze kilka minut wcześniej podobna
// myśl nie miała szans pojawić się w umyśle Białka.
// - Maciek, ty tak na serio? Oddasz mi się? - zapytał z niedowierzaniem. Kiner skinął głową i oswobodził się ze spodni,
// które aż do tej chwili pętały mu kostki. Odwrócił się i ułożył w pozycji tylnej, która byłaby dla większości mężczyzn
// upokarzająca. Ciasna dziurka programisty zachęcająco przezierała przez gąszcz delikatnych włosków, aż prosząc się o rozepchanie.
// Penis Białka zareagował na ten widok od razu, ciekawie wynurzając się z rozpiętych jeszcze przez Elfika spodni.
// - Raz się żyje - pomyślał Białek. Wyruchanie kolegi z pracy mogło się okazać bardziej męskie, niż przypuszczał.
// To on miał przecież być stroną dominującą.
// Białek pozbył się resztek odzienia i pewnym krokiem podszedł do wypiętego Kinera. Postanowił zaatakować znienacka.
// Napluł na rękę i jednym ruchem wbił naprężonego kutasa w miękki odbyt programisty. Sięgnął ręką do podbrzusza,
// by przekonać się, że i chuj Kinera był tak nabrzmiały, jakby miał eksplodować. Dodało mu to siły i pewności siebie.
// Poczuł sie atrakcyjny.
// Pierwsze ruchy były jeszcze dość powolne, lecz Kiner starał się przyjąć jak najwięcej pomimo słabego nawilżenia.
// Kolejne centymetry napiętej pały Białka znikały w czeluściach jego odbytu, a zwierzęce sapanie Kinera tylko podniecało
// właściciela dość dużej jak na polskie warunki kutangi. Białek złapał Kinera za biodra i przyciągnął do siebie.
// - Chciałeś rżnięcia? To masz.
// - wycedził i zaczął miarowo, całym ciężarem napierać na puszyste ciałko kolegi.
// - Pierdol mnie, Michał!
// - krzyknął resztkami sił Kiner, uginając się do samej podłogi, a odgłosy największego
// rżnięcia w historii firmy ocuciły omdlałego Elfika, leżącego od dwa metry dalej.
// Elfik wstała i zataczając się podeszła do sapiącego Białka.
// - Hej, chłopakii... - z ust dziewczyny wychodził pijacki bełkot,
// który trudno było rozszyfrować.
// - mogę się przyłączyć?
// - Wypierdalaj stąd! - wrzasnął Białek i odtrącił rękę Elfika, która zaczęła mierzwić jego włosy.
// Elfik prychnęła i odeszła na kilka kroków. Udawała obrażoną, choć ciekawie zerkała na Białka, który bez opamiętania
// pierdolił Kinera jak maszyna. Sam Kiner odwrócił się w jej stronę i złośliwie wystawił język na wierzch.
// Wygrał tę partię.
// Poszukiwania piersiówki w torebce w panterkę okazały się owocne. Elfik pociągnęła resztkę bimbru dla kurażu i postanowiła
// nie rezygnować z szansy zostania moderatorką swojego ulubionego serwisu. Stanęła na wysokości oczu Białka i ostentacyjnie
// rozpoczęła striptiz. Na pierwszy ogień poszła spódniczka. Dopiero w tym świetle znać było ślady spermy i wymiocin,
// które pokrywały jasny materiał. Elfik zaplątała się w bluzkę, lecz niezrażona tym faktem wciąż starała się wyglądać seksi.
// Będąc już w samej bieliźnie odwróciła się i wypięła prosto przed twarzą Kinera, który niewiele myśląc splunął prosto
// na naddarty materiał elfikowych majtek. Elfik wybuchnęła śmiechem. Po zdjęciu stanika ułożyła się na podłodze i uchyliła majtki,
// pokazując nieco zarośniętą cipkę o wyraźnie zarysowanych wargach.
// - Nudzi mi się. - powiedziała.
// - Długo jeszcze będziecie się pierdolić?. Białek jednak nie odpowiedział, zbyt zaaferowany stanem przedorgazmicznym, który sobie zafundował.
// Największa kurewna wykopu postanowiła zabawić się sama. Sięgnęła ręką po trzonek od łopaty, która była maskotką serwisu
// i zaczęła nim jeździć po wargach sromowych. Śluz gęsto skapywał na podłogę, a Elfik pociągała się za sutki,
// wijąc się jak piskorz po tanich panelach. Trzonek wszedł w luźną jamę Elfika jak w masło. Prawdopodobnie nie czuła
// niczego, a choć starała się zwrócić na siebie uwagę przez jęki i udawane podniecenie, Białek i Kiner byli zajęci sobą.
// Spojeni w jedność dochodzili właśnie razem, o czym obwieścił światu pierwotny ryk rudego programisty. Zmęczony m__b od
// razu wyszedł z Kinera, racząc się widokiem ciepłego jeszcze ciasteczka z kremem. Na pożegnanie przytulił się do mięciutkiego
// tyłeczka kolegi, który dostarczył mu wiele satysfakcji.
// I - jak podpowiadała mu intuicja - miał dostarczyć jeszcze nieraz.
// Kutas Białka pokryty był lekko kałem, lecz Kiner sprawnie sobie z tym poradził, naprędce zlizując brązową maź z mięknącego
// już chuja administratora wykopu.
// - Byłeś zajebisty, nikt mnie jeszcze tak nie jebał. - powtarzał zmęczonym głosem Kiner.
// - Zdejmij skarpetki, chcę ci podziękować jeszcze bardziej.
// Zdziwiony Białek zsunął białe stopki z nóg, po czym Kiner rzucił się do ssania dużego palca. Znudzona Elfik naprędce
// znalazła się obok niego, próbując zmieścić w ustach jeszcze więcej palców, by zyskać sobie przychylność Białka.
// Ten zaś był w siódmym niebie. Nie spieszyło mu się już do domu tak, jak kilka godzin wcześniej.
// Kiner wykorzystał rozmarzenie kolegi, by na koniec usiąść mu na twarzy i zmusić go do ssania swoich kulek.
// W oczach Białka pojawiły się pierwsze ślady przywiązania, co bardzo Maćka wzruszylo. Nie zepsuło tej chwili nawet
// faux-pas Elfika, która odepchnięta zapachem stóp Białka zwymiotowała na jego nogi. Życie w serwisie już nigdy nie miało być takie samo.
