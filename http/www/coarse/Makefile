# https://github.ngxson.com/wllama/examples/basic/
# https://github.com/ngxson/wllama

embeddings:
	@make wllama
	@make bert

wllama:
	curl -sL -o javascript/wllama.js https://github.ngxson.com/wllama/esm/index.js
	curl -sL -o wasm/wllama.wasm https://github.ngxson.com/wllama/esm/wasm/wllama.wasm

bert:
	curl -sL -o models/bert-bge-small/ggml-model-f16.gguf https://huggingface.co/ggml-org/models/resolve/main/bert-bge-small/ggml-model-f16.gguf
