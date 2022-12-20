
.PHONY: load_config_from_json override_from_env auto_reload_config 

load_config_from_json:
	cd load_config_from_json && go run main.go
override_from_env:
	cd override_from_env && go run main.go
auto_reload_config:
	cd auto_reload_config && go run main.go


