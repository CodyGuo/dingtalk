package robot

// Text text类型
type Text struct {
	// Content 消息内容
	Content string `json:"content"`
}

func (Text) GetType() string {
	return M_TEXT.String()
}

func (t *Text) GetContent() string {
	return t.Content
}
func (t *Text) SetContent(content string) {
	t.Content = content
}

// Link link类型
type Link struct {
	// Title 消息标题
	Title string `json:"title"`

	// Text 消息内容。如果太长只会部分展示
	Text string `json:"text,omitempty"`

	// MessageURL 点击消息跳转的URL
	MessageURL string `json:"messageUrl"`

	// PicURL 图片URL
	PicURL string `json:"picUrl,omitempty"`
}

func (Link) GetType() string {
	return M_LINK.String()
}

func (l *Link) GetMessageUrl() string {
	return l.MessageURL
}
func (l *Link) SetMessageUrl(messageURL string) {
	l.MessageURL = messageURL
}
func (l *Link) GetPicURL() string {
	return l.PicURL
}
func (l *Link) SetPicURL(picURL string) {
	l.PicURL = picURL
}
func (l *Link) GetText() string {
	return l.Text
}
func (l *Link) SetText(text string) {
	l.Text = text
}
func (l *Link) GetTitle() string {
	return l.Title
}
func (l *Link) SetTitle(title string) {
	l.Title = title
}

// Markdown markdown类型
type Markdown struct {
	// Title 首屏会话透出的展示内容
	Title string `json:"title"`

	// Text  markdown格式的消息
	Text string `json:"text"`
}

func (Markdown) GetType() string {
	return M_MARKDOWN.String()
}

func (m *Markdown) GetTitle() string {
	return m.Title
}
func (m *Markdown) SetTitle(title string) {
	m.Title = title
}

func (m *Markdown) GetText() string {
	return m.Text
}
func (m *Markdown) SetText(text string) {
	m.Text = text
}

// Btn 按钮的信息
type Btn struct {
	// Title 按钮方案
	Title string `json:"title"`

	// ActionURL 点击按钮触发的URL
	ActionURL string `json:"actionURL"`
}

func (b *Btn) GetTitle() string {
	return b.Title
}

func (b *Btn) SetTitle(title string) {
	b.Title = title
}

func (b *Btn) GetActionURL() string {
	return b.ActionURL
}

func (b *Btn) SetActionURL(actionURL string) {
	b.ActionURL = actionURL
}

// ActionCard actionCard类型
type ActionCard struct {
	// Title 首屏会话透出的展示内容
	Title string `json:"title"`

	// Text markdown格式的消息
	Text string `json:"text"`

	// SingleTitle 单个按钮的方案。(设置此项和singleURL后btns无效)
	SingleTitle string `json:"singleTitle"`

	// SingleURL 点击singleTitle按钮触发的URL
	SingleURL string `json:"singleURL"`

	// HideAvatar      0-正常发消息者头像,1-隐藏发消息者头像
	HideAvatar string `json:"hideAvatar,omitempty"`

	// Btns 按钮的信息：title-按钮方案，actionURL-点击按钮触发的URL
	Btns []Btn `json:"btns,omitempty"`

	// BtnOrientation   0-按钮竖直排列，1-按钮横向排列
	BtnOrientation string `json:"btnOrientation,omitempty"`
}

func (ActionCard) GetType() string {
	return M_ACTIONCARD.String()
}

func (a *ActionCard) GetBtnOrientation() string {
	return a.BtnOrientation
}
func (a *ActionCard) SetBtnOrientation(btnOrientation string) {
	a.BtnOrientation = btnOrientation
}
func (a *ActionCard) GetBtns() []Btn {
	return a.Btns
}
func (a *ActionCard) SetBtns(btns []Btn) {
	a.Btns = btns
}
func (a *ActionCard) GetHideAvatar() string {
	return a.HideAvatar
}
func (a *ActionCard) SetHideAvatar(hideAvatar string) {
	a.HideAvatar = hideAvatar
}
func (a *ActionCard) GetSingleTitle() string {
	return a.SingleTitle
}
func (a *ActionCard) SetSingleTitle(singleTitle string) {
	a.SingleTitle = singleTitle
}
func (a *ActionCard) GetSingleURL() string {
	return a.SingleURL
}
func (a *ActionCard) SetSingleURL(singleURL string) {
	a.SingleURL = singleURL
}
func (a *ActionCard) GetText() string {
	return a.Text
}
func (a *ActionCard) SetText(text string) {
	a.Text = text
}
func (a *ActionCard) GetTitle() string {
	return a.Title
}
func (a *ActionCard) SetTitle(title string) {
	a.Title = title
}

type FeedCardLink struct {
	//Title 单条信息文本
	Title string `json:"title"`

	// MessageURL 点击单条信息到跳转链接
	MessageURL string `json:"messageURL"`

	// PicURL 单条信息后面图片的URL
	PicURL string `json:"picURL"`
}

func (f *FeedCardLink) GetTitle() string {
	return f.Title
}

func (f *FeedCardLink) SetTitle(title string) {
	f.Title = title
}

func (f *FeedCardLink) GetMessageURL() string {
	return f.MessageURL
}

func (f *FeedCardLink) SetMessageURL(messageURL string) {
	f.MessageURL = messageURL
}

func (f *FeedCardLink) GetPicURL() string {
	return f.PicURL
}

func (f *FeedCardLink) SetPicURL(picURL string) {
	f.PicURL = picURL
}

// FeedCard 资讯类信息
type FeedCard struct {
	Links []FeedCardLink `json:"links"`
}

func (FeedCard) GetType() string {
	return M_FEEDCARD.String()
}

func (f *FeedCard) GetLinks() []FeedCardLink {
	return f.Links
}

func (f *FeedCard) SetLinks(links []FeedCardLink) {
	f.Links = links
}
