package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)


// 定义一个全局的验证器实例
var Validate *validator.Validate

func init() {
	// 启用了对结构体字段的必需性检查, 即 `validate:"required"` 标签
	Validate = validator.New(validator.WithRequiredStructEnabled())
}



// 将数据编码为 JSON 格式并写入 HTTP 响应中, 同时设置正确的响应头和状态码
func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}


func writeJSONError(w http.ResponseWriter, status int, message string) error {
	type envelope struct {
		Error string `json:"error"`
	}

	return writeJSON(w, status, &envelope{Error: message})
}


// 从 HTTP 请求中读取 JSON 数据, 并将其解码到指定的数据结构中, 同时限制请求体的大小, 并防止未知字段的出现
// 或者 parseJSON
func readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1_048_576

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}


func (h *Handler) jsonResponse(w http.ResponseWriter, status int, data any) error {
	// 包裹 payload, resp
	type envelope struct {
		Data any `json:"data"`

	}

	return writeJSON(w, status, &envelope{ Data: data})
}

func (h *Handler) jsonify(w http.ResponseWriter, status bool, message string, data any) error {
	type envelope struct {
		Status bool `json:"status"`
		Message string `json:"message"`
		Data any `json:"data,omitempty"`
		
	}

	return writeJSON(w, http.StatusOK, &envelope{ Status: status, Message: message, Data: data})
}


