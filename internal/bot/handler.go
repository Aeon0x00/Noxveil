package bot

import (
	"errors"
	"strconv"
	"strings"

	"log"

	"github.com/Aeon0x00/Noxveil/internal/repository"
	"github.com/Aeon0x00/Noxveil/internal/xui"
	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

type Handler struct {
	userRepo *repository.UserRepository
	subRepo	*repository.SubRepo
}

func NewHandler (userRepo *repository.UserRepository, subRepo *repository.SubRepo) *Handler {
	return &Handler{
		userRepo: userRepo,
		subRepo: subRepo,
	}
}

func (h *Handler) Start(ctx *th.Context, message telego.Message) error {
	tgID := message.From.ID
	userName := message.From.Username
	user,err := h.userRepo.GetByTelegramID(ctx,tgID)
	if err != nil {
		if errors.Is(err,pgx.ErrNoRows){
			user,err = h.userRepo.CreateUser(ctx, tgID)
		}
	}
	isActive,err := h.subRepo.CheckOnActive(ctx,user.ID)

	_, _ = ctx.Bot().SendRichMessage(ctx, &telego.SendRichMessageParams{
		ChatID:      tu.ID(message.Chat.ID),
		RichMessage: *renderMainScreen(userName,user.Balance,isActive),
	})

	return nil
}

func (h *Handler) ConfigurePlan(ctx *th.Context, query telego.CallbackQuery) error {
	_ = ctx.Bot().AnswerCallbackQuery(ctx, tu.CallbackQuery(query.ID))

	parts := strings.Split(query.Data, ":")

	CurrentGB, _ := strconv.Atoi(parts[1])
	CurrentDevice, _ := strconv.Atoi(parts[2])

	_, _ = ctx.Bot().EditMessageText(ctx, &telego.EditMessageTextParams{
		ChatID:      query.Message.GetChat().ChatID(),
		MessageID:   query.Message.GetMessageID(),
		RichMessage: renderPlan(CurrentGB, CurrentDevice),
	})
	return nil
}

func (h *Handler) BuyPlan(ctx *th.Context, query telego.CallbackQuery) error {
	_ = ctx.Bot().AnswerCallbackQuery(ctx, tu.CallbackQuery(query.ID))

	tgUserID := query.From.ID
	parts := strings.Split(query.Data, ":")

	totalGB, _ := strconv.Atoi(parts[1])
	limitDevice, _ := strconv.Atoi(parts[2])

	sub, err := xui.Conn(tgUserID, totalGB, limitDevice)
	
	if err != nil {
		log.Println("Ошибка создаине ключа подключения")
		_, _ = ctx.Bot().EditMessageText(ctx, &telego.EditMessageTextParams{
			ChatID:    query.Message.GetChat().ChatID(),
			MessageID: query.Message.GetMessageID(),
			RichMessage: telego.ToPtr(tu.RichMessage(
				tu.RichBlockSectionHeading(tu.RichTextPlain("❌ Ошибка сервера"), 2),
				tu.RichBlockParagraph(tu.RichTextBold(tu.RichTextPlain("Не удалось выпустить ключ. Пожалуйста, обратитесь в поддержку."))),
				tu.RichBlockButtons(
					tu.RichMessageButton(tu.RichTextPlain("Назад")).WithCallbackData("back_to_main"),
				),
			)),
		})} 

	user,_ := h.userRepo.GetByTelegramID(ctx,tgUserID)
	h.subRepo.CreateSub(ctx,sub.ID,user.ID,sub.SubID,limitDevice,totalGB)
	_, _ = ctx.Bot().EditMessageText(ctx, &telego.EditMessageTextParams{
		ChatID:    query.Message.GetChat().ChatID(),
		MessageID: query.Message.GetMessageID(),
		RichMessage: telego.ToPtr(tu.RichMessage(
			tu.RichBlockSectionHeading(tu.RichTextPlain("✅ Успешно! Ваша ссылка:"), 2),
			tu.RichBlockParagraph(tu.RichTextPlain(sub.Link)),
			tu.RichBlockButtons(
				tu.RichMessageButton(tu.RichTextPlain("Назад")).WithCallbackData("back_to_main"),
			),
		)),
	})
	return nil
}

func (h *Handler) BackToMain(ctx *th.Context, query telego.CallbackQuery) error{
	tgID := query.From.ID
	user,err := h.userRepo.GetByTelegramID(ctx,tgID)
	if err != nil {
		if errors.Is(err,pgx.ErrNoRows){
			user,err = h.userRepo.CreateUser(ctx, tgID)
		}
	}
	isActive,err := h.subRepo.CheckOnActive(ctx,user.ID)
	userName := query.From.Username

	_ = ctx.Bot().AnswerCallbackQuery(ctx, tu.CallbackQuery(query.ID))

	_,_ = ctx.Bot().EditMessageText(ctx, &telego.EditMessageTextParams{
		ChatID: query.Message.GetChat().ChatID(),
		MessageID: query.Message.GetMessageID(),
		RichMessage: renderMainScreen(userName,user.Balance,isActive),
	})
	return nil
}

func (h *Handler) Register(bh *th.BotHandler) {
	bh.HandleMessage(h.Start,th.CommandEqual("start"))

	bh.HandleCallbackQuery(h.ConfigurePlan,
		th.AnyCallbackQueryWithMessage(),
		th.CallbackDataPrefix("cfg:"),
	)

	bh.HandleCallbackQuery(h.BuyPlan,
		th.AnyCallbackQueryWithMessage(),
		th.CallbackDataPrefix("buy:"),
	)

	bh.HandleCallbackQuery(h.BackToMain,
		th.AnyCallbackQueryWithMessage(),
		th.CallbackDataPrefix("back_to_main"),
	)
	
	
}
//
// }, th.AnyCallbackQueryWithMessage(), th.CallbackDataPrefix("buy:"))
//
// bh.HandleCallbackQuery(func(ctx *th.Context, query telego.CallbackQuery) error {
//
// }, th.AnyCallbackQueryWithMessage(), th.CallbackDataPrefix("back_to_main"))
