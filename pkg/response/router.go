package response

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/open-runtimes/types-for-go/v4/openruntimes"
)

// RouteContext 包装 openruntimes.Context 并添加路由参数支持
type RouteContext struct {
	*openruntimes.Context
	RouteParams map[string]string
}

// Route 路由信息
type Route struct {
	pattern    string
	regex      *regexp.Regexp
	paramNames []string
	handler    func(c *RouteContext) openruntimes.Response
}

// Router 路由器
type Router struct {
	// method -> []Route
	routes map[string][]Route
}

func NewRouter() *Router {
	return &Router{routes: make(map[string][]Route)}
}

func NewRouteContext(ctx *openruntimes.Context) *RouteContext {
	return &RouteContext{
		Context:     ctx,
		RouteParams: make(map[string]string),
	}
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

// 获取请求体
func (c *RouteContext) GetBodyBinary() []byte {
	return c.Context.Req.BodyBinary()
}

// 获取请求体
func (c *RouteContext) GetJsonBody() map[string]interface{} {
	var result map[string]interface{}
	json.Unmarshal(c.Context.Req.BodyBinary(), &result)
	return result
}

// 获取请求体
func (c *RouteContext) GetBody(key string) string {
	bodyMap := c.GetJsonBody()
	return bodyMap[key].(string)
}

// 获取请求头
func (c *RouteContext) GetHeaders() map[string]string {
	return c.Context.Req.Headers
}

// 获取请求头
func (c *RouteContext) GetHeader(key string) string {
	return c.Context.Req.Headers[key]
}

// 获取请求方法
func (c *RouteContext) GetMethod() string {
	return c.Context.Req.Method
}

// 获取请求URL
func (c *RouteContext) GetUrl() string {
	return c.Context.Req.Url
}

// 获取请求端口
func (c *RouteContext) GetPort() int {
	return c.Context.Req.Port
}

// 获取请求协议
func (c *RouteContext) GetScheme() string {
	return c.Context.Req.Scheme
}

// 获取请求主机
func (c *RouteContext) GetHost() string {
	return c.Context.Req.Host
}

// 获取请求参数
func (c *RouteContext) GetQueryString() string {
	return c.Context.Req.QueryString
}

// 获取请求参数
func (c *RouteContext) GetQuery() map[string]string {
	return c.Context.Req.Query
}

