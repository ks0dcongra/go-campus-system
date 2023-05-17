## 輸入參數作為環境變量傳遞至 Drone Plugin 內部，需要加上 PLUGIN_ 前綴詞。
if [ -z ${PLUGIN_HELLO} ]; then
  PLUGIN_HELLO="default"
fi