package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// PagesClearCacheBuilder builder
//
// Allows to clear the cache of particular 'external' pages which may be attached to VK posts.
//
// https://vk.com/dev/pages.clearCache
type PagesClearCacheBuilder struct {
	api.Params
}

// NewPagesClearCacheBuilder func
func NewPagesClearCacheBuilder() *PagesClearCacheBuilder {
	return &PagesClearCacheBuilder{api.Params{}}
}

// URL Address of the page where you need to refesh the cached version
func (b *PagesClearCacheBuilder) URL(v string) {
	b.Params["url"] = v
}

// PagesGetBuilder builder
//
// Returns information about a wiki page.
//
// https://vk.com/dev/pages.get
type PagesGetBuilder struct {
	api.Params
}

// NewPagesGetBuilder func
func NewPagesGetBuilder() *PagesGetBuilder {
	return &PagesGetBuilder{api.Params{}}
}

// OwnerID Page owner ID.
func (b *PagesGetBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PageID Wiki page ID.
func (b *PagesGetBuilder) PageID(v int) {
	b.Params["page_id"] = v
}

// Global '1' — to return information about a global wiki page
func (b *PagesGetBuilder) Global(v bool) {
	b.Params["global"] = v
}

// SitePreview '1' — resulting wiki page is a preview for the attached link
func (b *PagesGetBuilder) SitePreview(v bool) {
	b.Params["site_preview"] = v
}

// Title Wiki page title.
func (b *PagesGetBuilder) Title(v string) {
	b.Params["title"] = v
}

// NeedSource parameter
func (b *PagesGetBuilder) NeedSource(v bool) {
	b.Params["need_source"] = v
}

// NeedHTML '1' — to return the page as HTML,
func (b *PagesGetBuilder) NeedHTML(v bool) {
	b.Params["need_html"] = v
}

// PagesGetHistoryBuilder builder
//
// Returns a list of all previous versions of a wiki page.
//
// https://vk.com/dev/pages.getHistory
type PagesGetHistoryBuilder struct {
	api.Params
}

// NewPagesGetHistoryBuilder func
func NewPagesGetHistoryBuilder() *PagesGetHistoryBuilder {
	return &PagesGetHistoryBuilder{api.Params{}}
}

// PageID Wiki page ID.
func (b *PagesGetHistoryBuilder) PageID(v int) {
	b.Params["page_id"] = v
}

// GroupID ID of the community that owns the wiki page.
func (b *PagesGetHistoryBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// UserID parameter
func (b *PagesGetHistoryBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// PagesGetTitlesBuilder builder
//
// Returns a list of wiki pages in a group.
//
// https://vk.com/dev/pages.getTitles
type PagesGetTitlesBuilder struct {
	api.Params
}

// NewPagesGetTitlesBuilder func
func NewPagesGetTitlesBuilder() *PagesGetTitlesBuilder {
	return &PagesGetTitlesBuilder{api.Params{}}
}

// GroupID ID of the community that owns the wiki page.
func (b *PagesGetTitlesBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// PagesGetVersionBuilder builder
//
// Returns the text of one of the previous versions of a wiki page.
//
// https://vk.com/dev/pages.getVersion
type PagesGetVersionBuilder struct {
	api.Params
}

// NewPagesGetVersionBuilder func
func NewPagesGetVersionBuilder() *PagesGetVersionBuilder {
	return &PagesGetVersionBuilder{api.Params{}}
}

// VersionID parameter
func (b *PagesGetVersionBuilder) VersionID(v int) {
	b.Params["version_id"] = v
}

// GroupID ID of the community that owns the wiki page.
func (b *PagesGetVersionBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// UserID parameter
func (b *PagesGetVersionBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// NeedHTML '1' — to return the page as HTML
func (b *PagesGetVersionBuilder) NeedHTML(v bool) {
	b.Params["need_html"] = v
}

// PagesParseWikiBuilder builder
//
// Returns HTML representation of the wiki markup.
//
// https://vk.com/dev/pages.parseWiki
type PagesParseWikiBuilder struct {
	api.Params
}

// NewPagesParseWikiBuilder func
func NewPagesParseWikiBuilder() *PagesParseWikiBuilder {
	return &PagesParseWikiBuilder{api.Params{}}
}

// Text Text of the wiki page.
func (b *PagesParseWikiBuilder) Text(v string) {
	b.Params["text"] = v
}

// GroupID ID of the group in the context of which this markup is interpreted.
func (b *PagesParseWikiBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// PagesSaveBuilder builder
//
// Saves the text of a wiki page.
//
// https://vk.com/dev/pages.save
type PagesSaveBuilder struct {
	api.Params
}

// NewPagesSaveBuilder func
func NewPagesSaveBuilder() *PagesSaveBuilder {
	return &PagesSaveBuilder{api.Params{}}
}

// Text Text of the wiki page in wiki-format.
func (b *PagesSaveBuilder) Text(v string) {
	b.Params["text"] = v
}

// PageID Wiki page ID. The 'title' parameter can be passed instead of 'pid'.
func (b *PagesSaveBuilder) PageID(v int) {
	b.Params["page_id"] = v
}

// GroupID ID of the community that owns the wiki page.
func (b *PagesSaveBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// UserID User ID
func (b *PagesSaveBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// Title Wiki page title.
func (b *PagesSaveBuilder) Title(v string) {
	b.Params["title"] = v
}

// PagesSaveAccessBuilder builder
//
// Saves modified read and edit access settings for a wiki page.
//
// https://vk.com/dev/pages.saveAccess
type PagesSaveAccessBuilder struct {
	api.Params
}

// NewPagesSaveAccessBuilder func
func NewPagesSaveAccessBuilder() *PagesSaveAccessBuilder {
	return &PagesSaveAccessBuilder{api.Params{}}
}

// PageID Wiki page ID.
func (b *PagesSaveAccessBuilder) PageID(v int) {
	b.Params["page_id"] = v
}

// GroupID ID of the community that owns the wiki page.
func (b *PagesSaveAccessBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// UserID parameter
func (b *PagesSaveAccessBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// View Who can view the wiki page:
//
// * 1 — only community members;
//
// * 2 — all users can view the page;
//
// * 0 — only community managers
func (b *PagesSaveAccessBuilder) View(v int) {
	b.Params["view"] = v
}

// Edit Who can edit the wiki page:
//
// * 1 — only community members;
//
// * 2 — all users can edit the page;
//
// * 0 — only community managers
func (b *PagesSaveAccessBuilder) Edit(v int) {
	b.Params["edit"] = v
}
