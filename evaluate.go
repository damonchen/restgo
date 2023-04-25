package restgo

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"go.uber.org/zap"
)

var (
	// ErrNotExistKey err not exist key
	ErrNotExistKey = errors.New("not exists key")

	logger *zap.Logger
)

func init() {
	logger, _ = zap.NewProduction()
}

// StageExecutor 每个文件一个执行器
type StageExecutor struct {
	logger *zap.Logger
	env    *Environment
	stage  *TestStage
	debug  bool
}

func (e StageExecutor) evaluate() error {
	defer logger.Sync()

	stages := e.stage.Stages
	for _, stage := range stages {
		err := e.evaluateStage(stage)
		if err != nil {
			return err
		}
	}

	return nil

}

func (e StageExecutor) evaluateStage(stage Stage) error {
	// 在执行前，先将需要替换的值替换掉
	req, err := e.newRequest(stage.Request)
	if err != nil {
		return err
	}

	client := NewClient(req)
	resp, err := client.Do()
	if err != nil {
		return err
	}

	// 1. 得到的内容和期盼的内容是否一致？
	// 2. 将需要保存的内容提取出来，保存到环境中
	if stage.Response.StatusCode != resp.StatusCode {
		return errors.New(fmt.Sprintf("expect status code %d, got %d", stage.Response.StatusCode, resp.StatusCode))
	}

	for key, value := range stage.Response.Headers {
		v := resp.Headers[key]
		if v != value {
			return errors.New(fmt.Sprintf("expect header %s with value %s, got %s",
				key, value, v))
		}
	}

	for key, value := range stage.Response.JSON {
		v := resp.JSON[key]
		if v != value {
			return errors.New(fmt.Sprintf("expect body %s with value %v, got %v",
				key, value, v))
		}
	}

	for key, value := range stage.Response.Save.JSON {
		keys := strings.Split(value, ".")

		var r interface{}
		data := resp.JSON
		for _, k := range keys {
			item := data[k]
			switch item.(type) {
			case string:
				r = item
			case map[string]interface{}:
				data = item.(map[string]interface{})
			}
		}
		e.env.Set(key, r)
	}

	return nil
}

func (e StageExecutor) newRequest(req *Request) (*Request, error) {
	url, err := e.replaceString(req.URL)
	if err != nil {
		return nil, err
	}

	headers, err := e.newMap(req.Headers)
	if err != nil {
		return nil, err
	}

	data, err := e.newMap(req.Data)
	if err != nil {
		return nil, err
	}

	json, err := e.newMapInterface(req.JSON)
	if err != nil {
		return nil, err
	}

	params, err := e.newMap(req.Params)
	if err != nil {
		return nil, err
	}

	return &Request{
		URL:     url,
		Method:  req.Method,
		JSON:    json,
		Params:  params,
		Data:    data,
		Headers: headers,
	}, nil
}

func (e StageExecutor) newMapInterface(items map[string]interface{}) (map[string]interface{}, error) {
	newItems := map[string]interface{}{}
	for key, value := range items {
		switch value.(type) {
		case string:
			v, err := e.replaceString(value.(string))
			if err != nil {
				return nil, err
			}
			newItems[key] = v
		default:
			newItems[key] = value
		}

	}
	return newItems, nil
}

func (e StageExecutor) newMap(items map[string]string) (map[string]string, error) {
	newItems := map[string]string{}
	for key, value := range items {
		v, err := e.replaceString(value)
		if err != nil {
			return nil, err
		}
		newItems[key] = v
	}
	return newItems, nil
}

func (e StageExecutor) replaceString(s string) (string, error) {
	// 值的结构：{key_name: }
	re := regexp.MustCompile("{key_name: (.*)}")
	matches := re.FindStringSubmatch(s)
	if len(matches) > 0 {
		keyName := matches[1]
		if !e.env.Has(keyName) {
			return "", fmt.Errorf("not exist key %s, matches %+v", s, matches)
		}
	}

	return s, nil
}

func (e StageExecutor) log(format string, args ...interface{}) {
	v := fmt.Sprintf(format, args...)
	if e.debug {
		e.logger.Debug(v)
	}
}

func Evaluate(stage *TestStage, debug bool) error {
	env := NewEnvironment()
	executor := StageExecutor{
		logger: logger,
		env:    env,
		stage:  stage,
		debug:  debug,
	}
	return executor.evaluate()
}
