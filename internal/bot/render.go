package bot

import (
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

type Toggle struct {
	On		string
	Off 	string
}

var (
	AccessStatus = Toggle{
		On:  "🟢 Доступ: Активен",
		Off: "🔴 Доступ: Не активен (требуется оплата)",
	}
	
	SelectedMark = Toggle{
		On:  "✅",
		Off: "",
	}
)

func (t Toggle) Get(cond bool) string {
	if cond {
		return t.On
	}
	return t.Off
}

func renderMainScreen (userName string, balance float64,isActive bool) *telego.InputRichMessage {
	fmt.Println(isActive)
	return telego.ToPtr(tu.RichMessage(
		tu.RichBlockSectionHeading(tu.RichTextBold(tu.RichTextPlain("⚡ Выберете действие")), 2),
		tu.RichBlockBlockQuotation(tu.RichBlockParagraph(tu.RichTextBold(tu.RichTextPlain(fmt.Sprintf("👤 Профиль: @%s \n 💰Баланс: %.2f ₽\n %s",userName,balance,AccessStatus.Get(isActive)))))),
		tu.RichBlockButtons(
			tu.RichMessageButton(tu.RichTextCode(tu.RichTextPlain("🚀 Купить подписку"))).WithCallbackData("cfg:150:1"),
		),
		tu.RichBlockButtons(
			tu.RichMessageButton(tu.RichTextPlain("👤 Профиль ")).WithCallbackData("plan_100"),
			tu.RichMessageButton(tu.RichTextPlain(" 💳 Баланс ")).WithCallbackData("plan_100"),

		),
		tu.RichBlockButtons(
			tu.RichMessageButton(tu.RichTextPlain(" 🌍 Серверы ")).WithCallbackData("plan_100"),
			tu.RichMessageButton(tu.RichTextPlain(" 💬 Поддержка ")).WithCallbackData("plan_100"),
		),
		tu.RichBlockDetails(tu.RichTextPlain("💡 Как подключиться за 1 минуту?"),tu.RichBlockParagraph(tu.RichTextPlain(""))),
	),
	)
}





func renderPlan (currentGB int, currentDevice int) *telego.InputRichMessage  {
	return telego.ToPtr(tu.RichMessage(
		tu.RichBlockSectionHeading( ( tu.RichTextBold( tu.RichTextPlain("Составьте свой план"))), 1),
		tu.RichBlockBlockQuotation(tu.RichBlockParagraph(tu.RichTextBold(tu.RichTextPlain(fmt.Sprintf("📊 Трафик %d Гб\n📱 Устройства: %d шт\n💰 Итого: ∞ ₽",currentGB,currentDevice))))),
		tu.RichBlockParagraph(tu.RichTextBold(tu.RichTextPlain("📊 Объём трафика в GB:"))),
		tu.RichBlockButtons(
			tu.RichMessageButton(tu.RichTextBold(tu.RichTextPlain(formatBtn("150", 150 == currentGB)))).WithCallbackData(fmt.Sprintf("cfg:%d:%d",150,currentDevice)),
			tu.RichMessageButton(tu.RichTextBold(tu.RichTextPlain(formatBtn("300", 300 == currentGB)))).WithCallbackData(fmt.Sprintf("cfg:%d:%d",300,currentDevice)),
			tu.RichMessageButton(tu.RichTextBold(tu.RichTextPlain(formatBtn("500", 500 == currentGB)))).WithCallbackData(fmt.Sprintf("cfg:%d:%d",500,currentDevice)),
			tu.RichMessageButton(tu.RichTextBold(tu.RichTextPlain(formatBtn("1000", 1000 == currentGB)))).WithCallbackData(fmt.Sprintf("cfg:%d:%d",1000,currentDevice)),
			tu.RichMessageButton(tu.RichTextBold(tu.RichTextPlain(formatBtn("2000", 2000 == currentGB)))).WithCallbackData(fmt.Sprintf("cfg:%d:%d",2000,currentDevice)),
			tu.RichMessageButton(tu.RichTextBold(tu.RichTextPlain(formatBtn("∞", 0 == currentGB)))).WithCallbackData(fmt.Sprintf("cfg:%d:%d",0,currentDevice)),
		),
		tu.RichBlockParagraph(tu.RichTextBold(tu.RichTextPlain("📱 Количество устройств:"))),
		tu.RichBlockButtons(
			tu.RichMessageButton(tu.RichTextBold(tu.RichTextPlain(formatBtn("1", 1 == currentDevice)))).WithCallbackData(fmt.Sprintf("cfg:%d:%d",currentGB,1)),
			tu.RichMessageButton(tu.RichTextBold(tu.RichTextPlain(formatBtn("3", 3 == currentDevice)))).WithCallbackData(fmt.Sprintf("cfg:%d:%d",currentGB,3)),
			tu.RichMessageButton(tu.RichTextBold(tu.RichTextPlain(formatBtn("5", 5 == currentDevice)))).WithCallbackData(fmt.Sprintf("cfg:%d:%d",currentGB,5)),
			tu.RichMessageButton(tu.RichTextBold(tu.RichTextPlain(formatBtn("10", 10 == currentDevice)))).WithCallbackData(fmt.Sprintf("cfg:%d:%d",currentGB,10)),
			tu.RichMessageButton(tu.RichTextBold(tu.RichTextPlain(formatBtn("20", 20 == currentDevice)))).WithCallbackData(fmt.Sprintf("cfg:%d:%d",currentGB,20)),
			tu.RichMessageButton(tu.RichTextBold(tu.RichTextPlain(formatBtn("∞", 0 == currentDevice)))).WithCallbackData(fmt.Sprintf("cfg:%d:%d",currentGB,0)),

		),
		tu.RichBlockButtons(
			tu.RichMessageButton(tu.RichTextCode(tu.RichTextPlain("🚀 Купить подписку"))).WithCallbackData(fmt.Sprintf("buy:%d:%d",currentGB,currentDevice)),
		),
		tu.RichBlockButtons(
			tu.RichMessageButton(tu.RichTextCode(tu.RichTextPlain("Назад"))).WithCallbackData("back_to_main"),
		),
	))
}

func formatBtn(label string, isActive bool) string {
	return label + SelectedMark.Get(isActive)
}