package telegram

const (
	CommandProcessingError = "Помилка обробки команди. @UQuark"
	CommandNotFoundError = "Команда не знайдена."
	SendError = "Виникла помилка при надсиланні повідомлення. @UQuark"
	GettingAdminsError = "Виникла помилка при отриманні списку адміністраторів. @UQuark"

	PingPong                 = "Pong."
	SetReadonlyNthMediaRule  = "Автоматично видано RO на %d годин(и/у) за п. 2 правил Барренсу. Для оскарження покличте живого адміністратора."
	SetReadonlyFailedAdmin   = "Ахтунг! Я не можу видати RO адміністратору. Покличте живого адміністратора для цього."
	SetReadonlyFailedUnknown = "Виникла помилка при спробі видати RO. @UQuark"
	MustBeAReply             = "Ахтунг! Для виконання команди потрібно відповісти на цільове повідомлення."
	AdminRequired            = "Ахтунг! Тільки адміністратор може виконувати цю команду."
	SetReadonlyHour          = "Користувачу видано RO на %d годин(и/у) за командою адміністратора."
	SetReadonlyDay = "Користувачу видано RO на %d днів(день) за командою адміністратора."
)