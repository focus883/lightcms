package handlers

import (
	"net/http"
	"strings"
)

func adminTemplateLanguage(r *http.Request) string {
	if r == nil {
		return "en"
	}

	candidates := make([]string, 0, 4)
	if lang := strings.TrimSpace(r.URL.Query().Get("lang")); lang != "" {
		candidates = append(candidates, lang)
	}
	if c, err := r.Cookie("lang"); err == nil && c != nil {
		if lang := strings.TrimSpace(c.Value); lang != "" {
			candidates = append(candidates, lang)
		}
	}
	if lang := strings.TrimSpace(r.Header.Get("Accept-Language")); lang != "" {
		candidates = append(candidates, lang)
	}

	for _, candidate := range candidates {
		lower := strings.ToLower(candidate)
		if strings.Contains(lower, "zh") {
			return "zh-CN"
		}
		if strings.HasPrefix(lower, "en") {
			return "en"
		}
	}
	return "en"
}

func translateAdminTemplateSource(source string, r *http.Request) string {
	if source == "" || adminTemplateLanguage(r) != "zh-CN" {
		return source
	}

	replacer := strings.NewReplacer(
		"Login - LightCMS", "登录 - LightCMS",
		"Content Management System", "内容管理系统",
		"Email", "邮箱",
		"Enter email", "输入邮箱",
		"Password", "密码",
		"Enter password", "输入密码",
		"Sign In", "登录",
		"Dashboard", "控制台",
		"Content Items", "内容项",
		"Templates", "模板",
		"Daily Active Users", "日活跃用户",
		"Monthly Active Users", "月活跃用户",
		"Pages Created Today", "今日新增页面",
		"Quick Actions", "快捷操作",
		"New Content", "新建内容",
		"New Template", "新建模板",
		"View Site", "查看站点",
		"Recent Content", "最近内容",
		"Requiring Approval", "待审批",
		"View All", "查看全部",
		"Submitted By", "提交人",
		"Submitted", "提交时间",
		"Recent Comments", "最近评论",
		"Published", "已发布",
		"Draft", "草稿",
		"View", "查看",
		"Edit", "编辑",
		"Delete", "删除",
		"Name", "名称",
		"Category", "分类",
		"Fields", "字段",
		"System", "系统",
		"Actions", "操作",
		"Yes", "是",
		"No", "否",
		"Go back", "返回",
		"Not permitted", "无权限",
		"Dashboard", "控制台",
		"Are you sure you want to delete this template?", "确定删除此模板吗？",
	)

	return replacer.Replace(source)
}
