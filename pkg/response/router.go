package response

import (
	"regexp"
	"strings"

	"github.com/open-runtimes/types-for-go/v4/openruntimes"
)

// RouteContext 包装 openruntimes.Context 并添加路由参数支持
type RouteContext struct {
	*openruntimes.Context
	RouteParams map[string]string
}

type Route struct {
	pattern    string
	regex      *regexp.Regexp
	paramNames []string
	handler    func(c *RouteContext) openruntimes.Response
}

type Router struct {
	// method -> []Route
	routes map[string][]Route
}

func NewRouter() *Router {
	return &Router{routes: make(map[string][]Route)}
}

func (r *Router) AddRoute(pattern string, method string, handler func(c *RouteContext) openruntimes.Response) {
	// 初始化方法映射（如果不存在）
	if _, exists := r.routes[method]; !exists {
		r.routes[method] = []Route{}
	}

	// 提取参数名称并创建正则表达式
	paramNames := []string{}
	regexPattern := pattern

	// 查找所有 {paramName} 模式并替换为正则表达式
	re := regexp.MustCompile(`\{([^/]+)\}`)
	matches := re.FindAllStringSubmatch(pattern, -1)

	for _, match := range matches {
		paramNames = append(paramNames, match[1])
		// 将 {paramName} 替换为捕获组
		regexPattern = strings.Replace(regexPattern, match[0], `([^/]+)`, 1)
	}

	// 将路径模式转换为正则表达式
	regexPattern = "^" + regexPattern + "$"
	regex := regexp.MustCompile(regexPattern)

	// 添加路由
	r.routes[method] = append(r.routes[method], Route{
		pattern:    pattern,
		regex:      regex,
		paramNames: paramNames,
		handler:    handler,
	})
}

func (r *Router) Handle(ctx *openruntimes.Context) openruntimes.Response {
	method := ctx.Req.Method
	path := ctx.Req.Path

	// 查找匹配的路由
	routes, methodExists := r.routes[method]
	if !methodExists {
		return NewStatusErrorResponse(ctx, 404)
	}

	for _, route := range routes {
		matches := route.regex.FindStringSubmatch(path)
		if matches != nil {
			// 找到匹配的路由
			// 创建路由上下文
			routeCtx := &RouteContext{
				Context:     ctx,
				RouteParams: make(map[string]string),
			}

			// 从第1个开始，因为第0个是整个匹配的字符串
			for i, name := range route.paramNames {
				if i+1 < len(matches) {
					routeCtx.RouteParams[name] = matches[i+1]
				}
			}

			// 调用处理函数
			return route.handler(routeCtx)
		}
	}

	// 没有找到匹配的路由
	return NewStatusErrorResponse(ctx, 404)
}

// GetParam 获取路由参数值的辅助方法
func (c *RouteContext) GetParam(name string) string {
	return c.RouteParams[name]
}
