var __defProp = Object.defineProperty;
var __defProps = Object.defineProperties;
var __getOwnPropDescs = Object.getOwnPropertyDescriptors;
var __getOwnPropSymbols = Object.getOwnPropertySymbols;
var __hasOwnProp = Object.prototype.hasOwnProperty;
var __propIsEnum = Object.prototype.propertyIsEnumerable;
var __knownSymbol = (name, symbol) => (symbol = Symbol[name]) ? symbol : Symbol.for("Symbol." + name);
var __defNormalProp = (obj, key, value) => key in obj ? __defProp(obj, key, { enumerable: true, configurable: true, writable: true, value }) : obj[key] = value;
var __spreadValues = (a, b) => {
  for (var prop in b || (b = {}))
    if (__hasOwnProp.call(b, prop))
      __defNormalProp(a, prop, b[prop]);
  if (__getOwnPropSymbols)
    for (var prop of __getOwnPropSymbols(b)) {
      if (__propIsEnum.call(b, prop))
        __defNormalProp(a, prop, b[prop]);
    }
  return a;
};
var __spreadProps = (a, b) => __defProps(a, __getOwnPropDescs(b));
var __publicField = (obj, key, value) => __defNormalProp(obj, typeof key !== "symbol" ? key + "" : key, value);
var __async = (__this, __arguments, generator) => {
  return new Promise((resolve, reject) => {
    var fulfilled = (value) => {
      try {
        step(generator.next(value));
      } catch (e) {
        reject(e);
      }
    };
    var rejected = (value) => {
      try {
        step(generator.throw(value));
      } catch (e) {
        reject(e);
      }
    };
    var step = (x) => x.done ? resolve(x.value) : Promise.resolve(x.value).then(fulfilled, rejected);
    step((generator = generator.apply(__this, __arguments)).next());
  });
};
var __await = function(promise, isYieldStar) {
  this[0] = promise;
  this[1] = isYieldStar;
};
var __asyncGenerator = (__this, __arguments, generator) => {
  var resume = (k, v, yes, no) => {
    try {
      var x = generator[k](v), isAwait = (v = x.value) instanceof __await, done = x.done;
      Promise.resolve(isAwait ? v[0] : v).then((y) => isAwait ? resume(k === "return" ? k : "next", v[1] ? { done: y.done, value: y.value } : y, yes, no) : yes({ value: y, done })).catch((e) => resume("throw", e, yes, no));
    } catch (e) {
      no(e);
    }
  }, method = (k) => it[k] = (x) => new Promise((yes, no) => resume(k, x, yes, no)), it = {};
  return generator = generator.apply(__this, __arguments), it[__knownSymbol("asyncIterator")] = () => it, method("next"), method("throw"), method("return"), it;
};
var __forAwait = (obj, it, method) => (it = obj[__knownSymbol("asyncIterator")]) ? it.call(obj) : (obj = obj[__knownSymbol("iterator")](), it = {}, method = (key, fn) => (fn = obj[key]) && (it[key] = (arg) => new Promise((yes, no, done) => (arg = fn.call(obj, arg), done = arg.done, Promise.resolve(arg.value).then((value) => yes({ value, done }), no)))), method("next"), method("return"), it);

// src/glue/messages.ts
var GLUE_VERSION = 1;
var GLUE_MESSAGE_PROTOTYPES = {
  "erro_evt": {
    "name": "erro_evt",
    "structName": "glue_msg_error",
    "className": "GlueMsgError",
    "fields": [
      {
        "type": "str",
        "name": "message",
        "isNullable": false
      }
    ]
  },
  "load_req": {
    "name": "load_req",
    "structName": "glue_msg_load_req",
    "className": "GlueMsgLoadReq",
    "fields": [
      {
        "type": "arr_str",
        "name": "model_paths",
        "isNullable": false
      },
      {
        "type": "str",
        "name": "mmproj_path",
        "isNullable": true
      },
      {
        "type": "bool",
        "name": "n_ctx_auto",
        "isNullable": false
      },
      {
        "type": "bool",
        "name": "use_mmap",
        "isNullable": false
      },
      {
        "type": "bool",
        "name": "use_mlock",
        "isNullable": false
      },
      {
        "type": "int",
        "name": "n_gpu_layers",
        "isNullable": false
      },
      {
        "type": "int",
        "name": "n_ctx",
        "isNullable": false
      },
      {
        "type": "int",
        "name": "n_threads",
        "isNullable": false
      },
      {
        "type": "str",
        "name": "model_alias",
        "isNullable": true
      },
      {
        "type": "int",
        "name": "log_level",
        "isNullable": true
      },
      {
        "type": "bool",
        "name": "embeddings",
        "isNullable": true
      },
      {
        "type": "bool",
        "name": "offload_kqv",
        "isNullable": true
      },
      {
        "type": "int",
        "name": "n_batch",
        "isNullable": true
      },
      {
        "type": "int",
        "name": "n_ubatch",
        "isNullable": true
      },
      {
        "type": "int",
        "name": "n_parallel",
        "isNullable": true
      },
      {
        "type": "str",
        "name": "pooling_type",
        "isNullable": true
      },
      {
        "type": "str",
        "name": "rope_scaling_type",
        "isNullable": true
      },
      {
        "type": "float",
        "name": "rope_freq_base",
        "isNullable": true
      },
      {
        "type": "float",
        "name": "rope_freq_scale",
        "isNullable": true
      },
      {
        "type": "float",
        "name": "yarn_ext_factor",
        "isNullable": true
      },
      {
        "type": "float",
        "name": "yarn_attn_factor",
        "isNullable": true
      },
      {
        "type": "float",
        "name": "yarn_beta_fast",
        "isNullable": true
      },
      {
        "type": "float",
        "name": "yarn_beta_slow",
        "isNullable": true
      },
      {
        "type": "int",
        "name": "yarn_orig_ctx",
        "isNullable": true
      },
      {
        "type": "str",
        "name": "cache_type_k",
        "isNullable": true
      },
      {
        "type": "str",
        "name": "cache_type_v",
        "isNullable": true
      },
      {
        "type": "bool",
        "name": "kv_unified",
        "isNullable": true
      },
      {
        "type": "bool",
        "name": "flash_attn",
        "isNullable": true
      },
      {
        "type": "bool",
        "name": "swa_full",
        "isNullable": true
      },
      {
        "type": "int",
        "name": "n_ctx_checkpoints",
        "isNullable": true
      },
      {
        "type": "int",
        "name": "checkpoint_min_step",
        "isNullable": true
      },
      {
        "type": "str",
        "name": "chat_template",
        "isNullable": true
      },
      {
        "type": "bool",
        "name": "jinja",
        "isNullable": true
      },
      {
        "type": "arr_str",
        "name": "default_template_kwargs_keys",
        "isNullable": true
      },
      {
        "type": "arr_str",
        "name": "default_template_kwargs_vals",
        "isNullable": true
      },
      {
        "type": "bool",
        "name": "reasoning",
        "isNullable": true
      },
      {
        "type": "int",
        "name": "image_min_tokens",
        "isNullable": true
      },
      {
        "type": "int",
        "name": "image_max_tokens",
        "isNullable": true
      },
      {
        "type": "bool",
        "name": "warmup",
        "isNullable": true
      },
      {
        "type": "bool",
        "name": "no_kv_offload",
        "isNullable": true
      },
      {
        "type": "bool",
        "name": "mmproj_offload",
        "isNullable": true
      },
      {
        "type": "bool",
        "name": "cont_batching",
        "isNullable": true
      },
      {
        "type": "int",
        "name": "n_keep",
        "isNullable": true
      },
      {
        "type": "bool",
        "name": "ctx_shift",
        "isNullable": true
      },
      {
        "type": "bool",
        "name": "cache_idle_slots",
        "isNullable": true
      },
      {
        "type": "int",
        "name": "n_cache_reuse",
        "isNullable": true
      },
      {
        "type": "arr_str",
        "name": "lora_paths",
        "isNullable": true
      },
      {
        "type": "arr_float",
        "name": "lora_scales",
        "isNullable": true
      },
      {
        "type": "bool",
        "name": "lora_init_without_apply",
        "isNullable": true
      },
      {
        "type": "str",
        "name": "spec_draft_model",
        "isNullable": true
      },
      {
        "type": "int",
        "name": "spec_draft_ngl",
        "isNullable": true
      },
      {
        "type": "int",
        "name": "spec_draft_n_max",
        "isNullable": true
      },
      {
        "type": "int",
        "name": "spec_draft_n_min",
        "isNullable": true
      },
      {
        "type": "float",
        "name": "spec_draft_p_min",
        "isNullable": true
      },
      {
        "type": "int",
        "name": "spec_draft_threads",
        "isNullable": true
      },
      {
        "type": "int",
        "name": "spec_draft_threads_batch",
        "isNullable": true
      },
      {
        "type": "arr_str",
        "name": "kv_overrides_keys",
        "isNullable": true
      },
      {
        "type": "arr_str",
        "name": "kv_overrides_vals",
        "isNullable": true
      },
      {
        "type": "int",
        "name": "reasoning_budget_tokens",
        "isNullable": true
      },
      {
        "type": "str",
        "name": "reasoning_budget_message",
        "isNullable": true
      },
      {
        "type": "str",
        "name": "reasoning_format",
        "isNullable": true
      },
      {
        "type": "bool",
        "name": "skip_chat_parsing",
        "isNullable": true
      },
      {
        "type": "bool",
        "name": "prefill_assistant",
        "isNullable": true
      }
    ]
  },
  "load_res": {
    "name": "load_res",
    "structName": "glue_msg_load_res",
    "className": "GlueMsgLoadRes",
    "fields": [
      {
        "type": "bool",
        "name": "success",
        "isNullable": false
      },
      {
        "type": "int",
        "name": "n_ctx",
        "isNullable": false
      },
      {
        "type": "int",
        "name": "n_batch",
        "isNullable": false
      },
      {
        "type": "int",
        "name": "n_ubatch",
        "isNullable": false
      },
      {
        "type": "int",
        "name": "n_vocab",
        "isNullable": false
      },
      {
        "type": "int",
        "name": "n_ctx_train",
        "isNullable": false
      },
      {
        "type": "int",
        "name": "n_embd",
        "isNullable": false
      },
      {
        "type": "int",
        "name": "n_layer",
        "isNullable": false
      },
      {
        "type": "arr_str",
        "name": "metadata_key",
        "isNullable": false
      },
      {
        "type": "arr_str",
        "name": "metadata_val",
        "isNullable": false
      },
      {
        "type": "int",
        "name": "token_bos",
        "isNullable": false
      },
      {
        "type": "int",
        "name": "token_eos",
        "isNullable": false
      },
      {
        "type": "int",
        "name": "token_eot",
        "isNullable": false
      },
      {
        "type": "arr_int",
        "name": "list_tokens_eog",
        "isNullable": false
      },
      {
        "type": "bool",
        "name": "add_bos_token",
        "isNullable": false
      },
      {
        "type": "bool",
        "name": "add_eos_token",
        "isNullable": false
      },
      {
        "type": "bool",
        "name": "has_encoder",
        "isNullable": false
      },
      {
        "type": "int",
        "name": "token_decoder_start",
        "isNullable": false
      },
      {
        "type": "str",
        "name": "media_marker",
        "isNullable": false
      },
      {
        "type": "bool",
        "name": "has_image_input",
        "isNullable": false
      },
      {
        "type": "bool",
        "name": "has_audio_input",
        "isNullable": false
      }
    ]
  },
  "cmpl_req": {
    "name": "cmpl_req",
    "structName": "glue_msg_completion_req",
    "className": "GlueMsgCompletionReq",
    "fields": [
      {
        "type": "bool",
        "name": "is_chat",
        "isNullable": false
      },
      {
        "type": "str",
        "name": "data_json",
        "isNullable": false
      },
      {
        "type": "arr_raw",
        "name": "files",
        "isNullable": false
      }
    ]
  },
  "cmpl_res": {
    "name": "cmpl_res",
    "structName": "glue_msg_completion_res",
    "className": "GlueMsgCompletionRes",
    "fields": [
      {
        "type": "bool",
        "name": "success",
        "isNullable": false
      },
      {
        "type": "int",
        "name": "req_id",
        "isNullable": false
      }
    ]
  },
  "embd_req": {
    "name": "embd_req",
    "structName": "glue_msg_embedding_req",
    "className": "GlueMsgEmbeddingReq",
    "fields": [
      {
        "type": "str",
        "name": "data_json",
        "isNullable": false
      },
      {
        "type": "arr_raw",
        "name": "files",
        "isNullable": false
      }
    ]
  },
  "embd_res": {
    "name": "embd_res",
    "structName": "glue_msg_embedding_res",
    "className": "GlueMsgEmbeddingRes",
    "fields": [
      {
        "type": "bool",
        "name": "success",
        "isNullable": false
      },
      {
        "type": "int",
        "name": "req_id",
        "isNullable": false
      }
    ]
  },
  "rrnk_req": {
    "name": "rrnk_req",
    "structName": "glue_msg_rerank_req",
    "className": "GlueMsgRerankReq",
    "fields": [
      {
        "type": "str",
        "name": "data_json",
        "isNullable": false
      }
    ]
  },
  "rrnk_res": {
    "name": "rrnk_res",
    "structName": "glue_msg_rerank_res",
    "className": "GlueMsgRerankRes",
    "fields": [
      {
        "type": "bool",
        "name": "success",
        "isNullable": false
      },
      {
        "type": "int",
        "name": "req_id",
        "isNullable": false
      }
    ]
  },
  "gres_req": {
    "name": "gres_req",
    "structName": "glue_msg_get_result_req",
    "className": "GlueMsgGetResultReq",
    "fields": [
      {
        "type": "int",
        "name": "req_id",
        "isNullable": false
      }
    ]
  },
  "gres_res": {
    "name": "gres_res",
    "structName": "glue_msg_get_result_res",
    "className": "GlueMsgGetResultRes",
    "fields": [
      {
        "type": "bool",
        "name": "success",
        "isNullable": false
      },
      {
        "type": "bool",
        "name": "has_more",
        "isNullable": false
      },
      {
        "type": "bool",
        "name": "is_error",
        "isNullable": false
      },
      {
        "type": "str",
        "name": "data_json",
        "isNullable": false
      }
    ]
  },
  "cncl_req": {
    "name": "cncl_req",
    "structName": "glue_msg_cancel_req",
    "className": "GlueMsgCancelReq",
    "fields": [
      {
        "type": "int",
        "name": "req_id",
        "isNullable": false
      }
    ]
  },
  "cncl_res": {
    "name": "cncl_res",
    "structName": "glue_msg_cancel_res",
    "className": "GlueMsgCancelRes",
    "fields": [
      {
        "type": "bool",
        "name": "success",
        "isNullable": false
      }
    ]
  },
  "tbop_req": {
    "name": "tbop_req",
    "structName": "glue_msg_test_backend_ops_req",
    "className": "GlueMsgTestBackendOpsReq",
    "fields": [
      {
        "type": "arr_str",
        "name": "args",
        "isNullable": false
      }
    ]
  },
  "tbop_res": {
    "name": "tbop_res",
    "structName": "glue_msg_test_backend_ops_res",
    "className": "GlueMsgTestBackendOpsRes",
    "fields": [
      {
        "type": "int",
        "name": "retcode",
        "isNullable": false
      },
      {
        "type": "bool",
        "name": "success",
        "isNullable": false
      }
    ]
  }
};

// src/glue/glue.ts
var GLUE_MAGIC = new Uint8Array([71, 76, 85, 69]);
var GLUE_DTYPE_NULL = 0;
var GLUE_DTYPE_BOOL = 1;
var GLUE_DTYPE_INT = 2;
var GLUE_DTYPE_FLOAT = 3;
var GLUE_DTYPE_STRING = 4;
var GLUE_DTYPE_RAW = 5;
var GLUE_DTYPE_ARRAY_BOOL = 6;
var GLUE_DTYPE_ARRAY_INT = 7;
var GLUE_DTYPE_ARRAY_FLOAT = 8;
var GLUE_DTYPE_ARRAY_STRING = 9;
var GLUE_DTYPE_ARRAY_RAW = 10;
var TYPE_MAP = {
  str: GLUE_DTYPE_STRING,
  int: GLUE_DTYPE_INT,
  float: GLUE_DTYPE_FLOAT,
  bool: GLUE_DTYPE_BOOL,
  raw: GLUE_DTYPE_RAW,
  arr_str: GLUE_DTYPE_ARRAY_STRING,
  arr_int: GLUE_DTYPE_ARRAY_INT,
  arr_float: GLUE_DTYPE_ARRAY_FLOAT,
  arr_bool: GLUE_DTYPE_ARRAY_BOOL,
  arr_raw: GLUE_DTYPE_ARRAY_RAW,
  null: GLUE_DTYPE_NULL
};
function glueDeserialize(buf) {
  let offset = 0;
  const view = new DataView(buf.buffer);
  const readUint32 = () => {
    const value = view.getUint32(offset, true);
    offset += 4;
    return value;
  };
  const readInt32 = () => {
    const value = view.getInt32(offset, true);
    offset += 4;
    return value;
  };
  const readFloat = () => {
    const value = view.getFloat32(offset, true);
    offset += 4;
    return value;
  };
  const readBool = () => {
    return readUint32() !== 0;
  };
  const readString = (customLen) => {
    const length = customLen != null ? customLen : readUint32();
    const value = new TextDecoder().decode(buf.slice(offset, offset + length));
    offset += length;
    return value;
  };
  const readRaw = () => {
    const length = readUint32();
    const value = buf.slice(offset, offset + length);
    offset += length;
    return value;
  };
  const readArray = (readItem) => {
    const length = readUint32();
    const value = new Array(length);
    for (let i = 0; i < length; i++) {
      value[i] = readItem();
    }
    return value;
  };
  const readNull = () => null;
  const readField = (field) => {
    switch (field.type) {
      case "str":
        return readString();
      case "int":
        return readInt32();
      case "float":
        return readFloat();
      case "bool":
        return readBool();
      case "raw":
        return readRaw();
      case "arr_str":
        return readArray(readString);
      case "arr_int":
        return readArray(readInt32);
      case "arr_float":
        return readArray(readFloat);
      case "arr_bool":
        return readArray(readBool);
      case "arr_raw":
        return readArray(readRaw);
      case "null":
        return readNull();
    }
  };
  const magicValid = buf[0] === GLUE_MAGIC[0] && buf[1] === GLUE_MAGIC[1] && buf[2] === GLUE_MAGIC[2] && buf[3] === GLUE_MAGIC[3];
  offset += 4;
  if (!magicValid) {
    throw new Error("Invalid magic number");
  }
  const version = readUint32();
  if (version !== GLUE_VERSION) {
    throw new Error("Invalid version number");
  }
  const name = readString(8);
  const msgProto = GLUE_MESSAGE_PROTOTYPES[name];
  if (!msgProto) {
    throw new Error(`Unknown message name: ${name}`);
  }
  const output = { _name: name };
  for (const field of msgProto.fields) {
    const readType = readUint32();
    if (readType === GLUE_DTYPE_NULL) {
      if (!field.isNullable) {
        throw new Error(
          `${name}: Expect field ${field.name} to be non-nullable`
        );
      }
      output[field.name] = null;
      continue;
    }
    if (readType !== TYPE_MAP[field.type]) {
      throw new Error(
        `${name}: Expect field ${field.name} to have type ${field.type}`
      );
    }
    output[field.name] = readField(field);
  }
  return output;
}
function glueSerialize(msg) {
  const msgProto = GLUE_MESSAGE_PROTOTYPES[msg._name];
  if (!msgProto) {
    throw new Error(`Unknown message name: ${msg._name}`);
  }
  const bufs = [];
  const writeUint32 = (value) => {
    const buf = new ArrayBuffer(4);
    new DataView(buf).setUint32(0, value, true);
    bufs.push(new Uint8Array(buf));
  };
  const writeInt32 = (value) => {
    const buf = new ArrayBuffer(4);
    new DataView(buf).setInt32(0, value, true);
    bufs.push(new Uint8Array(buf));
  };
  const writeFloat = (value) => {
    const buf = new ArrayBuffer(4);
    new DataView(buf).setFloat32(0, value, true);
    bufs.push(new Uint8Array(buf));
  };
  const writeBool = (value) => {
    writeUint32(value ? 1 : 0);
  };
  const writeString = (value) => {
    const utf8 = new TextEncoder().encode(value);
    writeUint32(utf8.byteLength);
    bufs.push(utf8);
  };
  const writeRaw = (value) => {
    writeUint32(value.byteLength);
    bufs.push(value);
  };
  const writeArray = (value, writeItem) => {
    writeUint32(value.length);
    for (const item of value) {
      writeItem(item);
    }
  };
  const writeNull = () => {
  };
  bufs.push(GLUE_MAGIC);
  writeUint32(GLUE_VERSION);
  {
    const utf8 = new TextEncoder().encode(msg._name);
    bufs.push(utf8);
  }
  for (const field of msgProto.fields) {
    const val = msg[field.name];
    if (!field.isNullable && (val === null || val === void 0)) {
      throw new Error(
        `${msg._name}: Expect field ${field.name} to be non-nullable`
      );
    }
    if (val === null || val === void 0) {
      writeUint32(GLUE_DTYPE_NULL);
      continue;
    }
    writeUint32(TYPE_MAP[field.type]);
    switch (field.type) {
      case "str":
        writeString(val);
        break;
      case "int":
        writeInt32(val);
        break;
      case "float":
        writeFloat(val);
        break;
      case "bool":
        writeBool(val);
        break;
      case "raw":
        writeRaw(val);
        break;
      case "arr_str":
        writeArray(val, writeString);
        break;
      case "arr_int":
        writeArray(val, writeInt32);
        break;
      case "arr_float":
        writeArray(val, writeFloat);
        break;
      case "arr_bool":
        writeArray(val, writeBool);
        break;
      case "arr_raw":
        writeArray(val, writeRaw);
        break;
      case "null":
        writeNull();
        break;
    }
  }
  const totalLength = bufs.reduce((acc, buf) => acc + buf.byteLength, 0);
  const output = new Uint8Array(totalLength);
  let offset = 0;
  for (const buf of bufs) {
    output.set(buf, offset);
    offset += buf.byteLength;
  }
  return output;
}

// src/wasm/source-map.ts
var WASM_SOURCE_MAP = {
  "default": "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",
  "compat": "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"
};

// src/debug.ts
var cache = /* @__PURE__ */ new Map();
function loadMap(buildKey) {
  return __async(this, null, function* () {
    if (cache.has(buildKey)) return cache.get(buildKey);
    const b64 = WASM_SOURCE_MAP[buildKey];
    if (!b64) throw new Error(`No source map for build "${buildKey}"`);
    const gzipped = Uint8Array.from(atob(b64), (c) => c.charCodeAt(0));
    const ds = new DecompressionStream("gzip");
    const writer = ds.writable.getWriter();
    writer.write(gzipped);
    writer.close();
    const buf = yield new Response(ds.readable).arrayBuffer();
    const dv = new DataView(buf);
    const bytes = new Uint8Array(buf);
    const firstId = dv.getUint32(0, true);
    const funcCount = dv.getUint32(4, true);
    const numNames = dv.getUint32(8, true);
    const td = new TextDecoder();
    const names = [];
    let pos = 12;
    for (let i = 0; i < numNames; i++) {
      const len = bytes[pos++];
      names.push(td.decode(bytes.subarray(pos, pos + len)));
      pos += len;
    }
    const funcNames = [];
    for (let i = 0; i < funcCount; i++) {
      const idx = dv.getUint16(pos, true);
      pos += 2;
      funcNames.push(idx === 65535 ? null : names[idx]);
    }
    const entry = { firstId, funcNames };
    cache.set(buildKey, entry);
    return entry;
  });
}
var Debug = {
  /**
   * Resolves a list of wasm function indices to their cleaned symbol names.
   */
  decodeFuncIds: (funcIds, isCompatBuild) => __async(void 0, null, function* () {
    const buildKey = isCompatBuild ? "compat" : "default";
    const { firstId, funcNames } = yield loadMap(buildKey);
    return funcIds.map((funcId) => {
      const i = funcId - firstId;
      const name = i >= 0 && i < funcNames.length && funcNames[i] ? funcNames[i] : "(unknown)";
      return { funcId, name };
    });
  }),
  /**
   * Annotates a wasm stack trace string with resolved function names.
   *
   * Example input from Chrome:
   *   at http://localhost:8080/esm/wasm/wllama.wasm:wasm-function[775]:0x74251
   *   at async blob:http://localhost:8080/53a863cc-7227-45cc-8594-ddbbf5257f20:317:28
   *
   * Example input from Firefox:
   *   @http://localhost:8080/esm/wasm/wllama.wasm:wasm-function[796]:0x7dfe2
   *       at wModuleInit/WebAssembly.promising/< (9b6a2acd-d909-44e2-b021-d42fb9087cfb:15:32) index.js:1433:45
   *
   * Example input from Safari:
   *   2441@wasm-function[2441]
   *       at wrapper (d746f19e-4523-4f36-ba06-d0969acc0b05:22:126009)
   *
   * Example output:
   *   wasm-func[775] (server_response::send)
   */
  decodeStackTrace: (stack, isCompatBuild) => __async(void 0, null, function* () {
    const re = /wasm-function\[(\d+)\]/g;
    const funcIds = [
      ...new Set([...stack.matchAll(re)].map((m) => parseInt(m[1])))
    ];
    if (funcIds.length === 0) return stack;
    const resolved = yield Debug.decodeFuncIds(funcIds, isCompatBuild);
    return resolved.map((r) => {
      if (r.name === "(unknown)") {
        return `    wasm-func[${r.funcId}] (unknown)`;
      }
      return `    wasm-func[${r.funcId}] (${r.name})`;
    }).join("\n");
  })
};

// src/utils.ts
var textDecoder = new TextDecoder();
var URL_PARTS_REGEX = /-(\d{5})-of-(\d{5})\.gguf(?:\?.*)?$/;
var parseShardNumber = (fnameOrUrl) => {
  const matches = fnameOrUrl.match(URL_PARTS_REGEX);
  if (!matches) {
    return {
      baseURL: fnameOrUrl,
      current: 1,
      total: 1
    };
  } else {
    return {
      baseURL: fnameOrUrl.replace(URL_PARTS_REGEX, ""),
      current: parseInt(matches[1]),
      total: parseInt(matches[2])
    };
  }
};
var sortFileByShard = (blobs) => {
  const isFiles = blobs.every((b) => !!b.name);
  if (isFiles && blobs.length > 1) {
    const files = blobs;
    files.sort((a, b) => {
      const infoA = parseShardNumber(a.name);
      const infoB = parseShardNumber(b.name);
      return infoA.current - infoB.current;
    });
  }
};
var isMmproj = (blob) => __async(void 0, null, function* () {
  const META_NAME = "general.architecture";
  const META_VAL = "clip";
  const tmp = blob.slice(0, 128 * 1024);
  const header = yield tmp.arrayBuffer();
  const buf = new Uint8Array(header);
  const nameBytes = new TextEncoder().encode(META_NAME);
  const valBytes = new TextEncoder().encode(META_VAL);
  let offset = -1;
  outer: for (let i = 0; i <= buf.length - nameBytes.length; i++) {
    for (let j = 0; j < nameBytes.length; j++) {
      if (buf[i + j] !== nameBytes[j]) continue outer;
    }
    offset = i;
    break;
  }
  if (offset === -1) return false;
  if (offset + 8 * 4 + 4 > buf.length) return false;
  const view = new DataView(header);
  const valLen = view.getBigUint64(offset + 8 * 3, true);
  if (valLen !== /* @__PURE__ */ BigInt("4")) return false;
  for (let i = 0; i < valBytes.length; i++) {
    if (buf[offset + 8 * 4 + i] !== valBytes[i]) return false;
  }
  return true;
});
var absoluteUrl = (relativePath) => typeof document === "undefined" ? new URL(relativePath, self.location.href).href : new URL(relativePath, document.baseURI).href;
var padDigits = (number, digits) => {
  return Array(Math.max(digits - String(number).length + 1, 0)).join("0") + number;
};
var sumArr = (arr) => arr.reduce((prev, curr) => prev + curr, 0);
var isString = (value) => !!(value == null ? void 0 : value.startsWith);
var MMPROJ_FILE_NAME = "mmproj.gguf";
var prepareBlobs = (blobsInp) => __async(void 0, null, function* () {
  const blobs = [];
  let blobMmproj = null;
  for (const blob of blobsInp) {
    if (yield isMmproj(blob)) {
      blobMmproj = blob;
    } else {
      blobs.push(blob);
    }
  }
  sortFileByShard(blobs);
  const result = blobs.map((blob, i) => ({
    blob,
    name: `model-${padDigits(i + 1, 5)}-of-${padDigits(blobs.length, 5)}.gguf`
  }));
  if (blobMmproj) {
    result.push({
      blob: blobMmproj,
      name: MMPROJ_FILE_NAME
    });
  }
  return {
    llm: result.filter((f) => f.name !== MMPROJ_FILE_NAME),
    mmproj: blobMmproj ? { blob: blobMmproj, name: MMPROJ_FILE_NAME } : null,
    all: result
  };
});
var isSupportMultiThread = () => ((e) => __async(void 0, null, function* () {
  try {
    return "undefined" != typeof MessageChannel && new MessageChannel().port1.postMessage(new SharedArrayBuffer(1)), WebAssembly.validate(e);
  } catch (e2) {
    return false;
  }
}))(
  new Uint8Array([
    0,
    97,
    115,
    109,
    1,
    0,
    0,
    0,
    1,
    4,
    1,
    96,
    0,
    0,
    3,
    2,
    1,
    0,
    5,
    4,
    1,
    3,
    1,
    1,
    10,
    11,
    1,
    9,
    0,
    65,
    0,
    254,
    16,
    2,
    0,
    26,
    11
  ])
);
var isSupportExceptions = () => __async(void 0, null, function* () {
  return WebAssembly.validate(
    new Uint8Array([
      0,
      97,
      115,
      109,
      1,
      0,
      0,
      0,
      1,
      4,
      1,
      96,
      0,
      0,
      3,
      2,
      1,
      0,
      10,
      8,
      1,
      6,
      0,
      6,
      64,
      25,
      11,
      11
    ])
  );
});
var isSupportSIMD = () => __async(void 0, null, function* () {
  return WebAssembly.validate(
    new Uint8Array([
      0,
      97,
      115,
      109,
      1,
      0,
      0,
      0,
      1,
      5,
      1,
      96,
      0,
      1,
      123,
      3,
      2,
      1,
      0,
      10,
      10,
      1,
      8,
      0,
      65,
      0,
      253,
      15,
      253,
      98,
      11
    ])
  );
});
var isSupportJSPI = () => {
  return !!WebAssembly.Suspending;
};
var isSupportWebGPU = () => {
  return !!navigator.gpu;
};
var isSupportMem64 = () => {
  try {
    new WebAssembly.Memory({
      address: "i64",
      initial: /* @__PURE__ */ BigInt("1")
      // 1 page (64 KiB)
    });
    return true;
  } catch (e) {
    return false;
  }
};
var checkEnvironmentCompatible = () => __async(void 0, null, function* () {
  if (!(yield isSupportExceptions())) {
    throw new Error("WebAssembly runtime does not support exception handling");
  }
  if (!(yield isSupportSIMD())) {
    throw new Error("WebAssembly runtime does not support SIMD");
  }
});
var isFirefox = () => {
  return !!navigator.userAgent.match(/Firefox\/([0-9\.]+)(?:\s|$)/);
};
var GGUF_FILE_REGEX = /^.*\.gguf(?:\?.*)?$/;
var isValidGgufFile = (path) => {
  return GGUF_FILE_REGEX.test(path);
};
var isSafariMobile = () => {
  return !!navigator.userAgent.match(/Version\/([0-9\._]+).*Mobile.*Safari.*/);
};
var createWorker = (workerCode) => {
  const workerURL = URL.createObjectURL(
    isString(workerCode) ? new Blob([workerCode], { type: "text/javascript" }) : workerCode
  );
  return new Worker(workerURL, { type: "module" });
};
var cbToAsyncIter = (fn) => (...args) => {
  let values = [];
  let resolve;
  let reject;
  values.push(
    new Promise((res, rej) => {
      resolve = res;
      reject = rej;
    })
  );
  fn(...args, (val, done, err) => {
    if (err) {
      reject(err);
      return;
    }
    resolve([val, done]);
    values.push(
      new Promise((res, rej) => {
        resolve = res;
        reject = rej;
      })
    );
  });
  return function() {
    return __asyncGenerator(this, null, function* () {
      let val;
      for (let i = 0, done = false; !done; i++) {
        [val, done] = yield new __await(values[i]);
        delete values[i];
        if (val !== void 0) yield val;
      }
    });
  }();
};
var canUseAsyncFileRead = (compat) => isSupportJSPI() || compat;
var needCompat = () => !isSupportJSPI() || !isSupportMem64();

// src/workers-code/generated.ts
var LIBLLAMA_VERSION = "b1-f16050d";
var LLAMA_CPP_WORKER_CODE = "// Start the main llama.cpp\nlet wllamaMalloc;\nlet wllamaStart;\nlet wllamaAction;\nlet wllamaExit;\nlet wllamaDebug;\n\nlet Module = null;\nlet isCompat = false;\nlet lastStack = '';\nlet isAborted = false;\nlet hasMultithread = false;\n\n//////////////////////////////////////////////////////////////\n// UTILS\n//////////////////////////////////////////////////////////////\n\n// send message back to main thread\nconst msg = (data, transfer) => postMessage(data, transfer);\n\n// Convert CPP log into JS log\nconst cppLogToJSLog = (line) => {\n  const matched = line.match(/@@(DEBUG|INFO|WARN|ERROR)@@(.*)/);\n  return !!matched\n    ? {\n        level: (matched[1] === 'INFO' ? 'debug' : matched[1]).toLowerCase(),\n        text: matched[2],\n      }\n    : { level: 'log', text: line };\n};\n\nconst getHeapU8 = () => {\n  const buffer = Module.wasmMemory.buffer;\n  return new Uint8Array(buffer);\n};\n\nconst toSizeT = (num) => {\n  return isCompat ? Number(num) : BigInt(num);\n};\n\n// Get module config that forwards stdout/err to main thread\nconst getWModuleConfig = (_argMainScriptBlob) => {\n  var pathConfig = RUN_OPTIONS.pathConfig;\n  var pthreadPoolSize = RUN_OPTIONS.nbThread;\n  var argMainScriptBlob = _argMainScriptBlob;\n\n  isCompat = RUN_OPTIONS.compat;\n  hasMultithread = pthreadPoolSize > 1;\n\n  msg({\n    verb: 'console.debug',\n    args: [\n      `Multithread enabled: ${hasMultithread}, pthreadPoolSize: ${pthreadPoolSize}`,\n    ],\n  });\n\n  if (!pathConfig['wllama.wasm']) {\n    throw new Error('\"wllama.wasm\" is missing in pathConfig');\n  }\n  return {\n    noInitialRun: true,\n    print: function (text) {\n      if (arguments.length > 1)\n        text = Array.prototype.slice.call(arguments).join(' ');\n      msg({ verb: 'console.log', args: [text] });\n    },\n    printErr: function (text) {\n      if (arguments.length > 1)\n        text = Array.prototype.slice.call(arguments).join(' ');\n      if (text.startsWith('@@STACK@@')) {\n        lastStack = text.slice('@@STACK@@'.length);\n        return;\n      }\n      const logLine = cppLogToJSLog(text);\n      msg({ verb: 'console.' + logLine.level, args: [logLine.text] });\n    },\n    locateFile: function (filename, basePath) {\n      const p = pathConfig[filename];\n      const truncate = (str) =>\n        str.length > 128 ? `${str.substr(0, 128)}...` : str;\n      if (filename.match(/wllama\\.worker\\.js/)) {\n        msg({\n          verb: 'console.error',\n          args: [\n            '\"wllama.worker.js\" is removed from v2.2.1. Hint: make sure to clear browser\\'s cache.',\n          ],\n        });\n      } else {\n        msg({\n          verb: 'console.debug',\n          args: [`Loading \"${filename}\" from \"${truncate(p)}\"`],\n        });\n        return p;\n      }\n    },\n    mainScriptUrlOrBlob: hasMultithread\n      ? argMainScriptBlob\n      : 'throw new Error(\"Multithreading is not enabled\")',\n    pthreadPoolSize: hasMultithread ? pthreadPoolSize : 0,\n    wasmMemory: hasMultithread ? getWasmMemory() : null,\n    onAbort: function (message) {\n      isAborted = true;\n      msg({ verb: 'signal.abort', args: ['abort', message, lastStack, null] });\n    },\n    onExit: function (code) {\n      isAborted = true;\n      const callstack = new Error().stack.toString();\n      msg({\n        verb: 'signal.abort',\n        args: ['abort', 'exit(' + code + ')', callstack, null],\n      });\n    },\n  };\n};\n\n// Get the memory to be used by wasm. (Only used in multi-thread mode)\n// Because we have a weird OOM issue on iOS, we need to try some values\n// See: https://github.com/emscripten-core/emscripten/issues/19144\n//      https://github.com/godotengine/godot/issues/70621\nconst getWasmMemory = () => {\n  let minBytes = 128 * 1024 * 1024;\n  let maxBytes = 4096 * 1024 * 1024;\n  let stepBytes = 128 * 1024 * 1024;\n  while (maxBytes > minBytes) {\n    try {\n      const wasmMemory = new WebAssembly.Memory({\n        initial: toSizeT(minBytes / 65536),\n        maximum: toSizeT(maxBytes / 65536),\n        shared: true,\n        address: isCompat ? undefined : 'i64',\n      });\n      return wasmMemory;\n    } catch (e) {\n      maxBytes -= stepBytes;\n      continue; // retry\n    }\n  }\n  throw new Error('Cannot allocate WebAssembly.Memory');\n};\n\n//////////////////////////////////////////////////////////////\n// HEAPFS PATCH\n//////////////////////////////////////////////////////////////\n\n/**\n * By default, emscripten uses memfs. The way it works is by\n * allocating new Uint8Array in javascript heap. This is not good\n * because it requires files to be copied to wasm heap each time\n * a file is read.\n *\n * HeapFS is an alternative, which resolves this problem by\n * allocating space for file directly inside wasm heap. This\n * allows us to mmap without doing any copy.\n *\n * For llama.cpp, this is great because we use MAP_SHARED\n *\n * Ref: https://github.com/ngxson/wllama/pull/39\n * Ref: https://github.com/emscripten-core/emscripten/blob/main/src/library_memfs.js\n *\n * Note 29/05/2024 @ngxson\n * Due to ftell() being limited to MAX_LONG, we cannot load files bigger than 2^31 bytes (or 2GB)\n * Ref: https://github.com/emscripten-core/emscripten/blob/main/system/lib/libc/musl/src/stdio/ftell.c\n */\n\nconst fsNameToFile = {}; // map Name => File\nconst fsIdToFile = {}; // map ID => File\nlet currFileId = 0;\n\n// Patch and redirect memfs calls to wllama\nconst patchHeapFS = () => {\n  const m = Module;\n  // save functions\n  m.MEMFS.stream_ops._read = m.MEMFS.stream_ops.read;\n  m.MEMFS.stream_ops._write = m.MEMFS.stream_ops.write;\n  m.MEMFS.stream_ops._llseek = m.MEMFS.stream_ops.llseek;\n  m.MEMFS.stream_ops._allocate = m.MEMFS.stream_ops.allocate;\n  m.MEMFS.stream_ops._mmap = m.MEMFS.stream_ops.mmap;\n  m.MEMFS.stream_ops._msync = m.MEMFS.stream_ops.msync;\n\n  const patchStream = (stream) => {\n    const name = stream.node.name;\n    if (fsNameToFile[name]) {\n      const f = fsNameToFile[name];\n      const ptr = Number(f.ptr);\n      stream.node.contents = getHeapU8().subarray(ptr, ptr + f.size);\n      stream.node.usedBytes = f.size;\n    }\n  };\n\n  // replace \"read\" functions\n  m.MEMFS.stream_ops.read = function (\n    stream,\n    buffer,\n    offset,\n    length,\n    position\n  ) {\n    patchStream(stream);\n    return m.MEMFS.stream_ops._read(stream, buffer, offset, length, position);\n  };\n  m.MEMFS.ops_table.file.stream.read = m.MEMFS.stream_ops.read;\n\n  // replace \"llseek\" functions\n  m.MEMFS.stream_ops.llseek = function (stream, offset, whence) {\n    patchStream(stream);\n    return m.MEMFS.stream_ops._llseek(stream, offset, whence);\n  };\n  m.MEMFS.ops_table.file.stream.llseek = m.MEMFS.stream_ops.llseek;\n\n  // replace \"mmap\" functions\n  m.MEMFS.stream_ops.mmap = function (stream, length, position, prot, flags) {\n    patchStream(stream);\n    const name = stream.node.name;\n    if (fsNameToFile[name]) {\n      const f = fsNameToFile[name];\n      const mmapPtr = f.ptr + toSizeT(position);\n      return {\n        ptr: mmapPtr,\n        allocated: false,\n      };\n    } else {\n      return m.MEMFS.stream_ops._mmap(stream, length, position, prot, flags);\n    }\n  };\n  m.MEMFS.ops_table.file.stream.mmap = m.MEMFS.stream_ops.mmap;\n\n  // mount FS\n  m.FS.mkdir('/models');\n  m.FS.mount(m.MEMFS, { root: '.' }, '/models');\n};\n\n// Allocate a new file in wllama heapfs, returns file ID\nconst heapfsAlloc = (name, size, allocBuffer) => {\n  if (size < 1) {\n    throw new Error('File size must be bigger than 0');\n  }\n  const m = Module;\n  const ptr = toSizeT(allocBuffer ? m.mmapAlloc(size) : 0);\n  const file = {\n    ptr: ptr,\n    size: size,\n    id: currFileId++,\n  };\n  fsIdToFile[file.id] = file;\n  fsNameToFile[name] = file;\n  return file.id;\n};\n\n// Add new file to wllama heapfs, return number of written bytes\nconst heapfsWrite = (id, buffer, offset) => {\n  if (fsIdToFile[id]) {\n    const { ptr, size } = fsIdToFile[id];\n    const afterWriteByte = offset + buffer.byteLength;\n    if (afterWriteByte > size) {\n      throw new Error(\n        `File ID ${id} write out of bound, afterWriteByte = ${afterWriteByte} while size = ${size}`\n      );\n    }\n    getHeapU8().set(buffer, Number(ptr) + offset);\n    return buffer.byteLength;\n  } else {\n    throw new Error(`File ID ${id} not found in heapfs`);\n  }\n};\n\n//////////////////////////////////////////////////////////////\n// ASYNC FILE READ\n//////////////////////////////////////////////////////////////\n\nlet isAwaitReading = false;\nlet pendingReadPromise = null;\nlet pendingReadResolve = null;\nlet pendingReadReject = null;\n\nconst _stripModelsPrefix = (path) => path.replace(/^\\/?models\\//, '');\n\n// Called from EM_ASYNC_JS stub in wllama-fs.h (path is already a JS string)\nconst _wllama_js_file_read = async (path, offset, req_size, out_ptr) => {\n  const name = _stripModelsPrefix(path);\n\n  pendingReadPromise = new Promise((res, rej) => {\n    pendingReadResolve = res;\n    pendingReadReject = rej;\n  });\n  isAwaitReading = true;\n\n  postMessage({ verb: 'fs.read_req', args: [name, offset, req_size] });\n\n  let data;\n  try {\n    data = await pendingReadPromise;\n  } finally {\n    isAwaitReading = false;\n    pendingReadResolve = null;\n    pendingReadReject = null;\n  }\n\n  const bytes = new Uint8Array(data);\n  getHeapU8().set(bytes, out_ptr);\n  return toSizeT(bytes.length);\n};\n\n//////////////////////////////////////////////////////////////\n// MAIN CODE\n//////////////////////////////////////////////////////////////\n\nconst callWrapper = (name, ret, args, isAsync) => {\n  const fn = Module.cwrap(\n    name,\n    ret,\n    args,\n    isAsync ? { async: true } : undefined\n  );\n  return async (action, req) => {\n    // console.log(`Calling ${name} with action:`, action, 'and req:', req);\n    let result;\n    try {\n      if (args.length === 2) {\n        result = isAsync ? await fn(action, req) : fn(action, req);\n      } else {\n        result = fn();\n      }\n    } catch (ex) {\n      console.error(ex);\n      throw ex;\n    }\n    return result;\n  };\n};\n\n// re-entering the wasm while a call is suspended (JSPI / asyncify) corrupts its state, so only one call runs at a time and the rest wait in the queue\nlet wasmCallBusy = false;\nconst wasmCallQueue = [];\n\nconst runWasmCall = async (callbackId, fn) => {\n  if (isAborted) {\n    // the wasm is dead, fail fast instead of calling into it\n    msg({ callbackId, err: 'wllama has crashed, please reload the module' });\n    return;\n  }\n  if (wasmCallBusy) {\n    wasmCallQueue.push({ callbackId, fn });\n    return;\n  }\n  wasmCallBusy = true;\n  try {\n    await fn();\n  } finally {\n    wasmCallBusy = false;\n    if (isAborted) {\n      // do not touch the wasm again after it aborted; the main thread already rejected the queued tasks\n      wasmCallQueue.length = 0;\n    } else {\n      const next = wasmCallQueue.shift();\n      if (next) runWasmCall(next.callbackId, next.fn);\n    }\n  }\n};\n\nconst runAction = async (data) => {\n  const { args, callbackId } = data;\n  const argAction = args[0];\n  const argEncodedMsg = args[1];\n  try {\n    const inputPtr = await wllamaMalloc(toSizeT(argEncodedMsg.byteLength), 0);\n    // copy data to wasm heap\n    const inputBuffer = new Uint8Array(\n      getHeapU8().buffer,\n      Number(inputPtr),\n      argEncodedMsg.byteLength\n    );\n    inputBuffer.set(argEncodedMsg, 0);\n    const outputPtr = await wllamaAction(argAction, inputPtr);\n    // length of output buffer is written at the first 4 bytes of input buffer\n    const outputLen = new Uint32Array(\n      getHeapU8().buffer,\n      Number(inputPtr),\n      1\n    )[0];\n    // copy the output buffer to JS heap\n    const outputBuffer = new Uint8Array(outputLen);\n    const outputSrcView = new Uint8Array(\n      getHeapU8().buffer,\n      Number(outputPtr),\n      outputLen\n    );\n    outputBuffer.set(outputSrcView, 0); // copy it\n    msg({ callbackId, result: outputBuffer }, [outputBuffer.buffer]);\n  } catch (err) {\n    handleError(err);\n  }\n};\n\nfunction handleError(err) {\n  // If WASM already aborted, onAbort already sent signal.abort; skip to avoid\n  // re-reporting the resulting WebAssembly.RuntimeError as a JS exception.\n  if (isAborted) return;\n\n  const message = err ? err.message || String(err) : 'Unknown error';\n  const stack = err ? err.stack || String(err) : '';\n  msg({\n    verb: 'signal.abort',\n    args: ['exception', message, stack, err],\n  });\n}\n\nonmessage = async (e) => {\n  if (!e.data) return;\n  const { verb, args, callbackId } = e.data;\n\n  // fs.read_res arrives while wasm is JSPI-suspended; resolve the pending promise.\n  if (verb === 'fs.read_res') {\n    if (pendingReadResolve) {\n      pendingReadResolve(args[0]);\n    }\n    return;\n  }\n\n  // Guard: while awaiting a file read, reject any other incoming task.\n  if (isAwaitReading) {\n    if (callbackId) {\n      msg({\n        callbackId,\n        err: 'Worker is suspended waiting for file data (JSPI)',\n      });\n    }\n    return;\n  }\n\n  if (!callbackId) {\n    msg({ verb: 'console.error', args: ['callbackId is required', e.data] });\n    return;\n  }\n\n  if (verb === 'module.init') {\n    const argMainScriptBlob = args[0];\n    const argUseAsyncFile = args[1];\n    try {\n      Module = getWModuleConfig(argMainScriptBlob);\n      Module.preRun = () => {\n        if (argUseAsyncFile) {\n          Module.ENV['USE_ASYNC_FILE'] = '1';\n        }\n      };\n      Module.onRuntimeInitialized = () => {\n        // async call once module is ready\n        // init FS\n        patchHeapFS();\n        // init cwrap\n        const pointer = isCompat ? 'number' : 'bigint';\n        // TODO: note sure why emscripten cannot bind if there is only 1 argument\n        wllamaMalloc = callWrapper('wllama_malloc', pointer, [\n          'number',\n          pointer,\n        ]);\n        wllamaStart = callWrapper('wllama_start', 'string', [], true);\n        wllamaAction = callWrapper(\n          'wllama_action',\n          pointer,\n          ['string', pointer],\n          true\n        );\n        wllamaExit = callWrapper('wllama_exit', 'string', []);\n        wllamaDebug = callWrapper('wllama_debug', 'string', []);\n        msg({ callbackId, result: null });\n      };\n      wModuleInit();\n    } catch (err) {\n      handleError(err);\n    }\n    return;\n  }\n\n  if (verb === 'fs.alloc') {\n    const argFilename = args[0];\n    const argSize = args[1];\n    const argAllocBuffer = args[2];\n    try {\n      // create blank file\n      const emptyBuffer = new ArrayBuffer(0);\n      Module['FS_createDataFile'](\n        '/models',\n        argFilename,\n        emptyBuffer,\n        true,\n        true,\n        true\n      );\n      // alloc data on heap\n      const fileId = heapfsAlloc(argFilename, argSize, argAllocBuffer);\n      msg({ callbackId, result: { fileId } });\n    } catch (err) {\n      handleError(err);\n    }\n    return;\n  }\n\n  if (verb === 'fs.write') {\n    const argFileId = args[0];\n    const argBuffer = args[1];\n    const argOffset = args[2];\n    try {\n      const writtenBytes = heapfsWrite(argFileId, argBuffer, argOffset);\n      msg({ callbackId, result: { writtenBytes } });\n    } catch (err) {\n      handleError(err);\n    }\n    return;\n  }\n\n  if (verb === 'wllama.start') {\n    await runWasmCall(callbackId, async () => {\n      try {\n        const result = await wllamaStart();\n        msg({ callbackId, result });\n      } catch (err) {\n        handleError(err);\n      }\n    });\n    return;\n  }\n\n  if (verb === 'wllama.action') {\n    await runWasmCall(callbackId, () => runAction(e.data));\n    return;\n  }\n\n  if (verb === 'wllama.exit') {\n    await runWasmCall(callbackId, async () => {\n      try {\n        const result = await wllamaExit();\n        msg({ callbackId, result });\n      } catch (err) {\n        handleError(err);\n      }\n    });\n    return;\n  }\n\n  if (verb === 'wllama.debug') {\n    await runWasmCall(callbackId, async () => {\n      try {\n        const result = await wllamaDebug();\n        msg({ callbackId, result });\n      } catch (err) {\n        handleError(err);\n      }\n    });\n    return;\n  }\n};\n";
var OPFS_UTILS_WORKER_CODE = "let accessHandle;\nlet abortController = new AbortController();\n\nasync function openFile(filename) {\n  const opfsRoot = await navigator.storage.getDirectory();\n  const cacheDir = await opfsRoot.getDirectoryHandle('cache', { create: true });\n  const fileHandler = await cacheDir.getFileHandle(filename, { create: true });\n  accessHandle = await fileHandler.createSyncAccessHandle();\n  accessHandle.truncate(0); // clear file content\n}\n\nasync function writeFile(buf) {\n  accessHandle.write(buf);\n}\n\nasync function closeFile() {\n  accessHandle.flush();\n  accessHandle.close();\n}\n\nasync function writeTextFile(filename, str) {\n  await openFile(filename);\n  await writeFile(new TextEncoder().encode(str));\n  await closeFile();\n}\n\nconst throttled = (func, delay) => {\n  let lastRun = 0;\n  return (...args) => {\n    const now = Date.now();\n    if (now - lastRun > delay) {\n      lastRun = now;\n      func.apply(null, args);\n    }\n  };\n};\n\nconst assertNonNull = (val) => {\n  if (val === null || val === undefined) {\n    throw new Error('OPFS Worker: Assertion failed');\n  }\n};\n\n// respond to main thread\nconst resOK = () => postMessage({ ok: true });\nconst resProgress = (loaded, total) =>\n  postMessage({ progress: { loaded, total } });\nconst resErr = (err) => postMessage({ err });\n\nonmessage = async (e) => {\n  try {\n    if (!e.data) return;\n\n    /**\n     * @param {Object} e.data\n     *\n     * Fine-control FS actions:\n     * - { action: 'open', filename: 'string' }\n     * - { action: 'write', buf: ArrayBuffer }\n     * - { action: 'close' }\n     *\n     * Simple write API:\n     * - { action: 'write-simple', filename: 'string', buf: ArrayBuffer }\n     *\n     * Download API:\n     * - { action: 'download', url: 'string', filename: 'string', options: Object, metadataFileName: 'string' }\n     * - { action: 'download-abort' }\n     */\n    const {\n      action,\n      filename,\n      buf,\n      url,\n      options,\n      metadataFileName,\n      metadataAdditional,\n    } = e.data;\n\n    if (action === 'open') {\n      assertNonNull(filename);\n      await openFile(filename);\n      return resOK();\n    } else if (action === 'write') {\n      assertNonNull(buf);\n      await writeFile(buf);\n      return resOK();\n    } else if (action === 'close') {\n      await closeFile();\n      return resOK();\n    } else if (action === 'write-simple') {\n      assertNonNull(filename);\n      assertNonNull(buf);\n      await openFile(filename);\n      await writeFile(buf);\n      await closeFile();\n      return resOK();\n    } else if (action === 'download') {\n      assertNonNull(url);\n      assertNonNull(filename);\n      assertNonNull(metadataFileName);\n      assertNonNull(options);\n      assertNonNull(options.aborted);\n      abortController = new AbortController();\n      if (options.aborted) abortController.abort();\n      const response = await fetch(url, {\n        ...options,\n        signal: abortController.signal,\n      });\n      const contentLength = response.headers.get('content-length');\n      const etag = (response.headers.get('etag') || '').replace(\n        /[^A-Za-z0-9]/g,\n        ''\n      );\n      const total = parseInt(contentLength, 10);\n      const reader = response.body.getReader();\n      await openFile(filename);\n      let loaded = 0;\n      const throttledProgress = throttled(resProgress, 100);\n      while (true) {\n        const { done, value } = await reader.read();\n        if (done) break;\n        loaded += value.byteLength;\n        await writeFile(value);\n        throttledProgress(loaded, total);\n      }\n      resProgress(total, total); // 100% done\n      await closeFile();\n      // make sure this is in-sync with CacheEntryMetadata\n      await writeTextFile(\n        metadataFileName,\n        JSON.stringify({\n          originalURL: url,\n          originalSize: total,\n          etag,\n          ...metadataAdditional,\n        })\n      );\n      return resOK();\n    } else if (action === 'download-abort') {\n      if (abortController) {\n        abortController.abort();\n      }\n      return;\n    }\n\n    throw new Error('OPFS Worker: Invalid action', e.data);\n  } catch (err) {\n    return resErr(err);\n  }\n};\n";
var WLLAMA_EMSCRIPTEN_CODE = 'var Module=typeof Module!="undefined"?Module:{};var ENVIRONMENT_IS_WEB=!!globalThis.window;var ENVIRONMENT_IS_WORKER=!!globalThis.WorkerGlobalScope;var ENVIRONMENT_IS_NODE=globalThis.process?.versions?.node&&globalThis.process?.type!="renderer";var ENVIRONMENT_IS_PTHREAD=ENVIRONMENT_IS_WORKER&&self.name?.startsWith("em-pthread");if(ENVIRONMENT_IS_NODE){var worker_threads=require("worker_threads");global.Worker=worker_threads.Worker;ENVIRONMENT_IS_WORKER=!worker_threads.isMainThread;ENVIRONMENT_IS_PTHREAD=ENVIRONMENT_IS_WORKER&&worker_threads["workerData"]=="em-pthread"}var arguments_=[];var thisProgram="./this.program";var quit_=(status,toThrow)=>{throw toThrow};var _scriptName=globalThis.document?.currentScript?.src;if(typeof __filename!="undefined"){_scriptName=__filename}else if(ENVIRONMENT_IS_WORKER){_scriptName=self.location.href}var scriptDirectory="";function locateFile(path){if(Module["locateFile"]){return Module["locateFile"](path,scriptDirectory)}return scriptDirectory+path}var readAsync,readBinary;if(ENVIRONMENT_IS_NODE){var fs=require("fs");scriptDirectory=__dirname+"/";readBinary=filename=>{filename=isFileURI(filename)?new URL(filename):filename;var ret=fs.readFileSync(filename);return ret};readAsync=async(filename,binary=true)=>{filename=isFileURI(filename)?new URL(filename):filename;var ret=fs.readFileSync(filename,binary?undefined:"utf8");return ret};if(process.argv.length>1){thisProgram=process.argv[1].replace(/\\\\/g,"/")}arguments_=process.argv.slice(2);if(typeof module!="undefined"){module["exports"]=Module}quit_=(status,toThrow)=>{process.exitCode=status;throw toThrow}}else if(ENVIRONMENT_IS_WEB||ENVIRONMENT_IS_WORKER){try{scriptDirectory=new URL(".",_scriptName).href}catch{}if(!ENVIRONMENT_IS_NODE){if(ENVIRONMENT_IS_WORKER){readBinary=url=>{var xhr=new XMLHttpRequest;xhr.open("GET",url,false);xhr.responseType="arraybuffer";xhr.send(null);return new Uint8Array(xhr.response)}}readAsync=async url=>{if(isFileURI(url)){return new Promise((resolve,reject)=>{var xhr=new XMLHttpRequest;xhr.open("GET",url,true);xhr.responseType="arraybuffer";xhr.onload=()=>{if(xhr.status==200||xhr.status==0&&xhr.response){resolve(xhr.response);return}reject(xhr.status)};xhr.onerror=reject;xhr.send(null)})}var response=await fetch(url,{credentials:"same-origin"});if(response.ok){return response.arrayBuffer()}throw new Error(response.status+" : "+response.url)}}}else{}var defaultPrint=console.log.bind(console);var defaultPrintErr=console.error.bind(console);if(ENVIRONMENT_IS_NODE){var utils=require("util");var stringify=a=>typeof a=="object"?utils.inspect(a):a;defaultPrint=(...args)=>fs.writeSync(1,args.map(stringify).join(" ")+"\\n");defaultPrintErr=(...args)=>fs.writeSync(2,args.map(stringify).join(" ")+"\\n")}var out=defaultPrint;var err=defaultPrintErr;var wasmBinary;var wasmModule;var ABORT=false;var EXITSTATUS;function assert(condition,text){if(!condition){abort(text)}}var isFileURI=filename=>filename.startsWith("file://");function growMemViews(){if(wasmMemory.buffer!=HEAP8.buffer){updateMemoryViews()}}if(ENVIRONMENT_IS_NODE&&ENVIRONMENT_IS_PTHREAD){var parentPort=worker_threads["parentPort"];parentPort.on("message",msg=>global.onmessage?.({data:msg}));Object.assign(globalThis,{self:global,postMessage:msg=>parentPort["postMessage"](msg)});process.on("uncaughtException",err=>{postMessage({cmd:"uncaughtException",error:err});process.exit(1)})}var startWorker;if(ENVIRONMENT_IS_PTHREAD){var initializedJS=false;self.onunhandledrejection=e=>{throw e.reason||e};async function handleMessage(e){try{var msgData=e["data"];var cmd=msgData.cmd;if(cmd==="load"){let messageQueue=[];self.onmessage=e=>messageQueue.push(e);startWorker=()=>{postMessage({cmd:"loaded"});for(let msg of messageQueue){handleMessage(msg)}self.onmessage=handleMessage};for(const handler of msgData.handlers){if(!Module[handler]||Module[handler].proxy){Module[handler]=(...args)=>{postMessage({cmd:"callHandler",handler,args})};if(handler=="print")out=Module[handler];if(handler=="printErr")err=Module[handler]}}wasmMemory=msgData.wasmMemory;updateMemoryViews();wasmModule=msgData.wasmModule;createWasm();run()}else if(cmd==="run"){establishStackSpace(msgData.pthread_ptr);__emscripten_thread_init(msgData.pthread_ptr,0,0,1,0,0);PThread.threadInitTLS();__emscripten_thread_mailbox_await(msgData.pthread_ptr);if(!initializedJS){initializedJS=true}try{await invokeEntryPoint(msgData.start_routine,msgData.arg)}catch(ex){if(ex!="unwind"){throw ex}}}else if(msgData.target==="setimmediate"){}else if(cmd==="checkMailbox"){if(initializedJS){checkMailbox()}}else if(cmd){err(`worker: received unknown command ${cmd}`);err(msgData)}}catch(ex){__emscripten_thread_crashed();throw ex}}self.onmessage=handleMessage}var HEAP8,HEAPU8,HEAP16,HEAPU16,HEAP32,HEAPU32,HEAPF32,HEAPF64;var HEAP64,HEAPU64;var runtimeInitialized=false;function updateMemoryViews(){var b=wasmMemory.buffer;HEAP8=new Int8Array(b);HEAP16=new Int16Array(b);Module["HEAPU8"]=HEAPU8=new Uint8Array(b);HEAPU16=new Uint16Array(b);HEAP32=new Int32Array(b);HEAPU32=new Uint32Array(b);HEAPF32=new Float32Array(b);HEAPF64=new Float64Array(b);HEAP64=new BigInt64Array(b);HEAPU64=new BigUint64Array(b)}function initMemory(){if(ENVIRONMENT_IS_PTHREAD){return}if(Module["wasmMemory"]){wasmMemory=Module["wasmMemory"]}else{var INITIAL_MEMORY=Module["INITIAL_MEMORY"]||134217728;wasmMemory=new WebAssembly.Memory({initial:BigInt(INITIAL_MEMORY/65536),maximum:65536n,shared:true,address:"i64"})}updateMemoryViews()}function preRun(){if(Module["preRun"]){if(typeof Module["preRun"]=="function")Module["preRun"]=[Module["preRun"]];while(Module["preRun"].length){addOnPreRun(Module["preRun"].shift())}}callRuntimeCallbacks(onPreRuns)}function initRuntime(){runtimeInitialized=true;if(ENVIRONMENT_IS_PTHREAD)return startWorker();if(!Module["noFSInit"]&&!FS.initialized)FS.init();TTY.init();wasmExports["__wasm_call_ctors"]();FS.ignorePermissions=false}function preMain(){}function postRun(){if(ENVIRONMENT_IS_PTHREAD){return}if(Module["postRun"]){if(typeof Module["postRun"]=="function")Module["postRun"]=[Module["postRun"]];while(Module["postRun"].length){addOnPostRun(Module["postRun"].shift())}}callRuntimeCallbacks(onPostRuns)}function abort(what){Module["onAbort"]?.(what);what="Aborted("+what+")";err(what);ABORT=true;what+=". Build with -sASSERTIONS for more info.";if(runtimeInitialized){___trap()}var e=new WebAssembly.RuntimeError(what);throw e}var wasmBinaryFile;function findWasmBinary(){return locateFile("wllama.wasm")}function getBinarySync(file){if(file==wasmBinaryFile&&wasmBinary){return new Uint8Array(wasmBinary)}if(readBinary){return readBinary(file)}throw"both async and sync fetching of the wasm failed"}async function getWasmBinary(binaryFile){if(!wasmBinary){try{var response=await readAsync(binaryFile);return new Uint8Array(response)}catch{}}return getBinarySync(binaryFile)}async function instantiateArrayBuffer(binaryFile,imports){try{var binary=await getWasmBinary(binaryFile);var instance=await WebAssembly.instantiate(binary,imports);return instance}catch(reason){err(`failed to asynchronously prepare wasm: ${reason}`);abort(reason)}}async function instantiateAsync(binary,binaryFile,imports){if(!binary&&!isFileURI(binaryFile)&&!ENVIRONMENT_IS_NODE){try{var response=fetch(binaryFile,{credentials:"same-origin"});var instantiationResult=await WebAssembly.instantiateStreaming(response,imports);return instantiationResult}catch(reason){err(`wasm streaming compile failed: ${reason}`);err("falling back to ArrayBuffer instantiation")}}return instantiateArrayBuffer(binaryFile,imports)}function getWasmImports(){assignWasmImports();if(!wasmImports.__instrumented){wasmImports.__instrumented=true;Asyncify.instrumentWasmImports(wasmImports)}var imports={env:wasmImports,wasi_snapshot_preview1:wasmImports};return imports}async function createWasm(){function receiveInstance(instance,module){wasmExports=instance.exports;wasmExports=Asyncify.instrumentWasmExports(wasmExports);wasmExports=applySignatureConversions(wasmExports);registerTLSInit(wasmExports["_emscripten_tls_init"]);assignWasmExports(wasmExports);wasmModule=module;removeRunDependency("wasm-instantiate");return wasmExports}addRunDependency("wasm-instantiate");function receiveInstantiationResult(result){return receiveInstance(result["instance"],result["module"])}var info=getWasmImports();if(Module["instantiateWasm"]){return new Promise((resolve,reject)=>{Module["instantiateWasm"](info,(inst,mod)=>{resolve(receiveInstance(inst,mod))})})}if(ENVIRONMENT_IS_PTHREAD){var instance=new WebAssembly.Instance(wasmModule,getWasmImports());return receiveInstance(instance,wasmModule)}wasmBinaryFile??=findWasmBinary();var result=await instantiateAsync(wasmBinary,wasmBinaryFile,info);var exports=receiveInstantiationResult(result);return exports}class ExitStatus{name="ExitStatus";constructor(status){this.message=`Program terminated with exit(${status})`;this.status=status}}var terminateWorker=worker=>{worker.terminate();worker.onmessage=e=>{}};var cleanupThread=pthread_ptr=>{var worker=PThread.pthreads[pthread_ptr];PThread.returnWorkerToPool(worker)};var callRuntimeCallbacks=callbacks=>{while(callbacks.length>0){callbacks.shift()(Module)}};var onPreRuns=[];var addOnPreRun=cb=>onPreRuns.push(cb);var runDependencies=0;var dependenciesFulfilled=null;var removeRunDependency=id=>{runDependencies--;Module["monitorRunDependencies"]?.(runDependencies);if(runDependencies==0){if(dependenciesFulfilled){var callback=dependenciesFulfilled;dependenciesFulfilled=null;callback()}}};var addRunDependency=id=>{runDependencies++;Module["monitorRunDependencies"]?.(runDependencies)};var spawnThread=threadParams=>{var worker=PThread.getNewWorker();if(!worker){return 6}PThread.runningWorkers.push(worker);PThread.pthreads[threadParams.pthread_ptr]=worker;worker.pthread_ptr=threadParams.pthread_ptr;var msg={cmd:"run",start_routine:threadParams.startRoutine,arg:threadParams.arg,pthread_ptr:threadParams.pthread_ptr};if(ENVIRONMENT_IS_NODE){worker.unref()}worker.postMessage(msg,threadParams.transferList);return 0};var runtimeKeepaliveCounter=0;var keepRuntimeAlive=()=>noExitRuntime||runtimeKeepaliveCounter>0;var stackSave=()=>_emscripten_stack_get_current();var stackRestore=val=>__emscripten_stack_restore(val);var stackAlloc=sz=>__emscripten_stack_alloc(sz);var proxyToMainThread=(funcIndex,emAsmAddr,sync,...callArgs)=>{var serializedNumCallArgs=callArgs.length*2;var sp=stackSave();var args=stackAlloc(serializedNumCallArgs*8);var b=args/8;for(var i=0;i<callArgs.length;i++){var arg=callArgs[i];if(typeof arg=="bigint"){(growMemViews(),HEAP64)[b+2*i]=1n;(growMemViews(),HEAP64)[b+2*i+1]=arg}else{(growMemViews(),HEAP64)[b+2*i]=0n;(growMemViews(),HEAPF64)[b+2*i+1]=arg}}var rtn=__emscripten_run_js_on_main_thread(funcIndex,emAsmAddr,serializedNumCallArgs,args,sync);stackRestore(sp);return rtn};function _proc_exit(code){if(ENVIRONMENT_IS_PTHREAD)return proxyToMainThread(0,0,1,code);EXITSTATUS=code;if(!keepRuntimeAlive()){PThread.terminateAllThreads();Module["onExit"]?.(code);ABORT=true}quit_(code,new ExitStatus(code))}function exitOnMainThread(returnCode){if(ENVIRONMENT_IS_PTHREAD)return proxyToMainThread(1,0,0,returnCode);_exit(returnCode)}var exitJS=(status,implicit)=>{EXITSTATUS=status;if(ENVIRONMENT_IS_PTHREAD){exitOnMainThread(status);throw"unwind"}_proc_exit(status)};var _exit=exitJS;var PThread={unusedWorkers:[],runningWorkers:[],tlsInitFunctions:[],pthreads:{},init(){if(!ENVIRONMENT_IS_PTHREAD){PThread.initMainThread()}},initMainThread(){var pthreadPoolSize=Module["pthreadPoolSize"];while(pthreadPoolSize--){PThread.allocateUnusedWorker()}addOnPreRun(async()=>{var pthreadPoolReady=PThread.loadWasmModuleToAllWorkers();addRunDependency("loading-workers");await pthreadPoolReady;removeRunDependency("loading-workers")})},terminateAllThreads:()=>{for(var worker of PThread.runningWorkers){terminateWorker(worker)}for(var worker of PThread.unusedWorkers){terminateWorker(worker)}PThread.unusedWorkers=[];PThread.runningWorkers=[];PThread.pthreads={}},returnWorkerToPool:worker=>{var pthread_ptr=worker.pthread_ptr;delete PThread.pthreads[pthread_ptr];PThread.unusedWorkers.push(worker);PThread.runningWorkers.splice(PThread.runningWorkers.indexOf(worker),1);worker.pthread_ptr=0;__emscripten_thread_free_data(pthread_ptr)},threadInitTLS(){PThread.tlsInitFunctions.forEach(f=>f())},loadWasmModuleToWorker:worker=>new Promise(onFinishedLoading=>{worker.onmessage=e=>{var d=e["data"];var cmd=d.cmd;if(d.targetThread&&d.targetThread!=_pthread_self()){var targetWorker=PThread.pthreads[d.targetThread];if(targetWorker){targetWorker.postMessage(d,d.transferList)}else{err(`Internal error! Worker sent a message "${cmd}" to target pthread ${d.targetThread}, but that thread no longer exists!`)}return}if(cmd==="checkMailbox"){checkMailbox()}else if(cmd==="spawnThread"){spawnThread(d)}else if(cmd==="cleanupThread"){callUserCallback(()=>cleanupThread(d.thread))}else if(cmd==="loaded"){worker.loaded=true;if(ENVIRONMENT_IS_NODE&&!worker.pthread_ptr){worker.unref()}onFinishedLoading(worker)}else if(d.target==="setimmediate"){worker.postMessage(d)}else if(cmd==="uncaughtException"){worker.onerror(d.error)}else if(cmd==="callHandler"){Module[d.handler](...d.args)}else if(cmd){err(`worker sent an unknown command ${cmd}`)}};worker.onerror=e=>{var message="worker sent an error!";err(`${message} ${e.filename}:${e.lineno}: ${e.message}`);throw e};if(ENVIRONMENT_IS_NODE){worker.on("message",data=>worker.onmessage({data}));worker.on("error",e=>worker.onerror(e))}var handlers=[];var knownHandlers=["onExit","onAbort","print","printErr"];for(var handler of knownHandlers){if(Module.propertyIsEnumerable(handler)){handlers.push(handler)}}worker.postMessage({cmd:"load",handlers,wasmMemory,wasmModule})}),async loadWasmModuleToAllWorkers(){if(ENVIRONMENT_IS_PTHREAD){return}let pthreadPoolReady=Promise.all(PThread.unusedWorkers.map(PThread.loadWasmModuleToWorker));return pthreadPoolReady},allocateUnusedWorker(){var worker;var pthreadMainJs=_scriptName;if(Module["mainScriptUrlOrBlob"]){pthreadMainJs=Module["mainScriptUrlOrBlob"];if(typeof pthreadMainJs!="string"){pthreadMainJs=URL.createObjectURL(pthreadMainJs)}}worker=new Worker(pthreadMainJs,{workerData:"em-pthread",name:"em-pthread"});PThread.unusedWorkers.push(worker)},getNewWorker(){if(PThread.unusedWorkers.length==0){PThread.allocateUnusedWorker();PThread.loadWasmModuleToWorker(PThread.unusedWorkers[0])}return PThread.unusedWorkers.pop()}};var onPostRuns=[];var addOnPostRun=cb=>onPostRuns.push(cb);function establishStackSpace(pthread_ptr){var stackHigh=Number((growMemViews(),HEAPU64)[(pthread_ptr+88)/8]);var stackSize=Number((growMemViews(),HEAPU64)[(pthread_ptr+96)/8]);var stackLow=stackHigh-stackSize;_emscripten_stack_set_limits(stackHigh,stackLow);stackRestore(stackHigh)}var wasmTableMirror=[];var getWasmTableEntry=funcPtr=>{funcPtr=Number(funcPtr);var func=wasmTableMirror[funcPtr];if(!func){wasmTableMirror[funcPtr]=func=wasmTable.get(BigInt(funcPtr));if(Asyncify.isAsyncExport(func)){wasmTableMirror[funcPtr]=func=Asyncify.makeAsyncFunction(func)}}return func};var invokeEntryPoint=async(ptr,arg)=>{runtimeKeepaliveCounter=0;noExitRuntime=0;var result=(a1=>WebAssembly.promising(getWasmTableEntry(ptr)).call(null,BigInt(a1)))(arg);function finish(result){if(keepRuntimeAlive()){EXITSTATUS=result;return}__emscripten_thread_exit(result)}result=await result;finish(result)};invokeEntryPoint.isAsync=true;var noExitRuntime=true;var registerTLSInit=tlsInitFunc=>PThread.tlsInitFunctions.push(tlsInitFunc);var wasmMemory;function pthreadCreateProxied(pthread_ptr,attr,startRoutine,arg){if(ENVIRONMENT_IS_PTHREAD)return proxyToMainThread(2,0,1,pthread_ptr,attr,startRoutine,arg);return ___pthread_create_js(pthread_ptr,attr,startRoutine,arg)}var _emscripten_has_threading_support=()=>!!globalThis.SharedArrayBuffer;var INT53_MAX=9007199254740992;var INT53_MIN=-9007199254740992;var bigintToI53Checked=num=>num<INT53_MIN||num>INT53_MAX?NaN:Number(num);function ___pthread_create_js(pthread_ptr,attr,startRoutine,arg){pthread_ptr=bigintToI53Checked(pthread_ptr);attr=bigintToI53Checked(attr);startRoutine=bigintToI53Checked(startRoutine);arg=bigintToI53Checked(arg);if(!_emscripten_has_threading_support()){return 6}var transferList=[];var error=0;if(ENVIRONMENT_IS_PTHREAD&&(transferList.length===0||error)){return pthreadCreateProxied(pthread_ptr,attr,startRoutine,arg)}if(error)return error;var threadParams={startRoutine,pthread_ptr,arg,transferList};if(ENVIRONMENT_IS_PTHREAD){threadParams.cmd="spawnThread";postMessage(threadParams,transferList);return 0}return spawnThread(threadParams)}var syscallGetVarargP=()=>{var ret=Number((growMemViews(),HEAPU64)[SYSCALLS.varargs/8]);SYSCALLS.varargs+=8;return ret};var syscallGetVarargI=()=>{var ret=(growMemViews(),HEAP32)[+SYSCALLS.varargs/4];SYSCALLS.varargs+=4;return ret};var PATH={isAbs:path=>path.charAt(0)==="/",splitPath:filename=>{var splitPathRe=/^(\\/?|)([\\s\\S]*?)((?:\\.{1,2}|[^\\/]+?|)(\\.[^.\\/]*|))(?:[\\/]*)$/;return splitPathRe.exec(filename).slice(1)},normalizeArray:(parts,allowAboveRoot)=>{var up=0;for(var i=parts.length-1;i>=0;i--){var last=parts[i];if(last==="."){parts.splice(i,1)}else if(last===".."){parts.splice(i,1);up++}else if(up){parts.splice(i,1);up--}}if(allowAboveRoot){for(;up;up--){parts.unshift("..")}}return parts},normalize:path=>{var isAbsolute=PATH.isAbs(path),trailingSlash=path.slice(-1)==="/";path=PATH.normalizeArray(path.split("/").filter(p=>!!p),!isAbsolute).join("/");if(!path&&!isAbsolute){path="."}if(path&&trailingSlash){path+="/"}return(isAbsolute?"/":"")+path},dirname:path=>{var result=PATH.splitPath(path),root=result[0],dir=result[1];if(!root&&!dir){return"."}if(dir){dir=dir.slice(0,-1)}return root+dir},basename:path=>path&&path.match(/([^\\/]+|\\/)\\/*$/)[1],join:(...paths)=>PATH.normalize(paths.join("/")),join2:(l,r)=>PATH.normalize(l+"/"+r)};var initRandomFill=()=>view=>view.set(crypto.getRandomValues(new Uint8Array(view.byteLength)));var randomFill=view=>{(randomFill=initRandomFill())(view)};var PATH_FS={resolve:(...args)=>{var resolvedPath="",resolvedAbsolute=false;for(var i=args.length-1;i>=-1&&!resolvedAbsolute;i--){var path=i>=0?args[i]:FS.cwd();if(typeof path!="string"){throw new TypeError("Arguments to path.resolve must be strings")}else if(!path){return""}resolvedPath=path+"/"+resolvedPath;resolvedAbsolute=PATH.isAbs(path)}resolvedPath=PATH.normalizeArray(resolvedPath.split("/").filter(p=>!!p),!resolvedAbsolute).join("/");return(resolvedAbsolute?"/":"")+resolvedPath||"."},relative:(from,to)=>{from=PATH_FS.resolve(from).slice(1);to=PATH_FS.resolve(to).slice(1);function trim(arr){var start=0;for(;start<arr.length;start++){if(arr[start]!=="")break}var end=arr.length-1;for(;end>=0;end--){if(arr[end]!=="")break}if(start>end)return[];return arr.slice(start,end-start+1)}var fromParts=trim(from.split("/"));var toParts=trim(to.split("/"));var length=Math.min(fromParts.length,toParts.length);var samePartsLength=length;for(var i=0;i<length;i++){if(fromParts[i]!==toParts[i]){samePartsLength=i;break}}var outputParts=[];for(var i=samePartsLength;i<fromParts.length;i++){outputParts.push("..")}outputParts=outputParts.concat(toParts.slice(samePartsLength));return outputParts.join("/")}};var UTF8Decoder=globalThis.TextDecoder&&new TextDecoder;var findStringEnd=(heapOrArray,idx,maxBytesToRead,ignoreNul)=>{var maxIdx=idx+maxBytesToRead;if(ignoreNul)return maxIdx;while(heapOrArray[idx]&&!(idx>=maxIdx))++idx;return idx};var UTF8ArrayToString=(heapOrArray,idx=0,maxBytesToRead,ignoreNul)=>{var endPtr=findStringEnd(heapOrArray,idx,maxBytesToRead,ignoreNul);if(endPtr-idx>16&&heapOrArray.buffer&&UTF8Decoder){return UTF8Decoder.decode(heapOrArray.buffer instanceof ArrayBuffer?heapOrArray.subarray(idx,endPtr):heapOrArray.slice(idx,endPtr))}var str="";while(idx<endPtr){var u0=heapOrArray[idx++];if(!(u0&128)){str+=String.fromCharCode(u0);continue}var u1=heapOrArray[idx++]&63;if((u0&224)==192){str+=String.fromCharCode((u0&31)<<6|u1);continue}var u2=heapOrArray[idx++]&63;if((u0&240)==224){u0=(u0&15)<<12|u1<<6|u2}else{u0=(u0&7)<<18|u1<<12|u2<<6|heapOrArray[idx++]&63}if(u0<65536){str+=String.fromCharCode(u0)}else{var ch=u0-65536;str+=String.fromCharCode(55296|ch>>10,56320|ch&1023)}}return str};var FS_stdin_getChar_buffer=[];var lengthBytesUTF8=str=>{var len=0;for(var i=0;i<str.length;++i){var c=str.charCodeAt(i);if(c<=127){len++}else if(c<=2047){len+=2}else if(c>=55296&&c<=57343){len+=4;++i}else{len+=3}}return len};var stringToUTF8Array=(str,heap,outIdx,maxBytesToWrite)=>{if(!(maxBytesToWrite>0))return 0;var startIdx=outIdx;var endIdx=outIdx+maxBytesToWrite-1;for(var i=0;i<str.length;++i){var u=str.codePointAt(i);if(u<=127){if(outIdx>=endIdx)break;heap[outIdx++]=u}else if(u<=2047){if(outIdx+1>=endIdx)break;heap[outIdx++]=192|u>>6;heap[outIdx++]=128|u&63}else if(u<=65535){if(outIdx+2>=endIdx)break;heap[outIdx++]=224|u>>12;heap[outIdx++]=128|u>>6&63;heap[outIdx++]=128|u&63}else{if(outIdx+3>=endIdx)break;heap[outIdx++]=240|u>>18;heap[outIdx++]=128|u>>12&63;heap[outIdx++]=128|u>>6&63;heap[outIdx++]=128|u&63;i++}}heap[outIdx]=0;return outIdx-startIdx};var intArrayFromString=(stringy,dontAddNull,length)=>{var len=length>0?length:lengthBytesUTF8(stringy)+1;var u8array=new Array(len);var numBytesWritten=stringToUTF8Array(stringy,u8array,0,u8array.length);if(dontAddNull)u8array.length=numBytesWritten;return u8array};var FS_stdin_getChar=()=>{if(!FS_stdin_getChar_buffer.length){var result=null;if(ENVIRONMENT_IS_NODE){var BUFSIZE=256;var buf=Buffer.alloc(BUFSIZE);var bytesRead=0;var fd=process.stdin.fd;try{bytesRead=fs.readSync(fd,buf,0,BUFSIZE)}catch(e){if(e.toString().includes("EOF"))bytesRead=0;else throw e}if(bytesRead>0){result=buf.slice(0,bytesRead).toString("utf-8")}}else if(globalThis.window?.prompt){result=window.prompt("Input: ");if(result!==null){result+="\\n"}}else{}if(!result){return null}FS_stdin_getChar_buffer=intArrayFromString(result,true)}return FS_stdin_getChar_buffer.shift()};var TTY={ttys:[],init(){},shutdown(){},register(dev,ops){TTY.ttys[dev]={input:[],output:[],ops};FS.registerDevice(dev,TTY.stream_ops)},stream_ops:{open(stream){var tty=TTY.ttys[stream.node.rdev];if(!tty){throw new FS.ErrnoError(43)}stream.tty=tty;stream.seekable=false},close(stream){stream.tty.ops.fsync(stream.tty)},fsync(stream){stream.tty.ops.fsync(stream.tty)},read(stream,buffer,offset,length,pos){if(!stream.tty||!stream.tty.ops.get_char){throw new FS.ErrnoError(60)}var bytesRead=0;for(var i=0;i<length;i++){var result;try{result=stream.tty.ops.get_char(stream.tty)}catch(e){throw new FS.ErrnoError(29)}if(result===undefined&&bytesRead===0){throw new FS.ErrnoError(6)}if(result===null||result===undefined)break;bytesRead++;buffer[offset+i]=result}if(bytesRead){stream.node.atime=Date.now()}return bytesRead},write(stream,buffer,offset,length,pos){if(!stream.tty||!stream.tty.ops.put_char){throw new FS.ErrnoError(60)}try{for(var i=0;i<length;i++){stream.tty.ops.put_char(stream.tty,buffer[offset+i])}}catch(e){throw new FS.ErrnoError(29)}if(length){stream.node.mtime=stream.node.ctime=Date.now()}return i}},default_tty_ops:{get_char(tty){return FS_stdin_getChar()},put_char(tty,val){if(val===null||val===10){out(UTF8ArrayToString(tty.output));tty.output=[]}else{if(val!=0)tty.output.push(val)}},fsync(tty){if(tty.output?.length>0){out(UTF8ArrayToString(tty.output));tty.output=[]}},ioctl_tcgets(tty){return{c_iflag:25856,c_oflag:5,c_cflag:191,c_lflag:35387,c_cc:[3,28,127,21,4,0,1,0,17,19,26,0,18,15,23,22,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]}},ioctl_tcsets(tty,optional_actions,data){return 0},ioctl_tiocgwinsz(tty){return[24,80]}},default_tty1_ops:{put_char(tty,val){if(val===null||val===10){err(UTF8ArrayToString(tty.output));tty.output=[]}else{if(val!=0)tty.output.push(val)}},fsync(tty){if(tty.output?.length>0){err(UTF8ArrayToString(tty.output));tty.output=[]}}}};var zeroMemory=(ptr,size)=>(growMemViews(),HEAPU8).fill(0,ptr,ptr+size);var alignMemory=(size,alignment)=>Math.ceil(size/alignment)*alignment;var mmapAlloc=size=>{size=alignMemory(size,65536);var ptr=_emscripten_builtin_memalign(65536,size);if(ptr)zeroMemory(ptr,size);return ptr};var MEMFS={ops_table:null,mount(mount){return MEMFS.createNode(null,"/",16895,0)},createNode(parent,name,mode,dev){if(FS.isBlkdev(mode)||FS.isFIFO(mode)){throw new FS.ErrnoError(63)}MEMFS.ops_table||={dir:{node:{getattr:MEMFS.node_ops.getattr,setattr:MEMFS.node_ops.setattr,lookup:MEMFS.node_ops.lookup,mknod:MEMFS.node_ops.mknod,rename:MEMFS.node_ops.rename,unlink:MEMFS.node_ops.unlink,rmdir:MEMFS.node_ops.rmdir,readdir:MEMFS.node_ops.readdir,symlink:MEMFS.node_ops.symlink},stream:{llseek:MEMFS.stream_ops.llseek}},file:{node:{getattr:MEMFS.node_ops.getattr,setattr:MEMFS.node_ops.setattr},stream:{llseek:MEMFS.stream_ops.llseek,read:MEMFS.stream_ops.read,write:MEMFS.stream_ops.write,mmap:MEMFS.stream_ops.mmap,msync:MEMFS.stream_ops.msync}},link:{node:{getattr:MEMFS.node_ops.getattr,setattr:MEMFS.node_ops.setattr,readlink:MEMFS.node_ops.readlink},stream:{}},chrdev:{node:{getattr:MEMFS.node_ops.getattr,setattr:MEMFS.node_ops.setattr},stream:FS.chrdev_stream_ops}};var node=FS.createNode(parent,name,mode,dev);if(FS.isDir(node.mode)){node.node_ops=MEMFS.ops_table.dir.node;node.stream_ops=MEMFS.ops_table.dir.stream;node.contents={}}else if(FS.isFile(node.mode)){node.node_ops=MEMFS.ops_table.file.node;node.stream_ops=MEMFS.ops_table.file.stream;node.usedBytes=0;node.contents=null}else if(FS.isLink(node.mode)){node.node_ops=MEMFS.ops_table.link.node;node.stream_ops=MEMFS.ops_table.link.stream}else if(FS.isChrdev(node.mode)){node.node_ops=MEMFS.ops_table.chrdev.node;node.stream_ops=MEMFS.ops_table.chrdev.stream}node.atime=node.mtime=node.ctime=Date.now();if(parent){parent.contents[name]=node;parent.atime=parent.mtime=parent.ctime=node.atime}return node},getFileDataAsTypedArray(node){if(!node.contents)return new Uint8Array(0);if(node.contents.subarray)return node.contents.subarray(0,node.usedBytes);return new Uint8Array(node.contents)},expandFileStorage(node,newCapacity){var prevCapacity=node.contents?node.contents.length:0;if(prevCapacity>=newCapacity)return;var CAPACITY_DOUBLING_MAX=1024*1024;newCapacity=Math.max(newCapacity,prevCapacity*(prevCapacity<CAPACITY_DOUBLING_MAX?2:1.125)>>>0);if(prevCapacity!=0)newCapacity=Math.max(newCapacity,256);var oldContents=node.contents;node.contents=new Uint8Array(newCapacity);if(node.usedBytes>0)node.contents.set(oldContents.subarray(0,node.usedBytes),0)},resizeFileStorage(node,newSize){if(node.usedBytes==newSize)return;if(newSize==0){node.contents=null;node.usedBytes=0}else{var oldContents=node.contents;node.contents=new Uint8Array(newSize);if(oldContents){node.contents.set(oldContents.subarray(0,Math.min(newSize,node.usedBytes)))}node.usedBytes=newSize}},node_ops:{getattr(node){var attr={};attr.dev=FS.isChrdev(node.mode)?node.id:1;attr.ino=node.id;attr.mode=node.mode;attr.nlink=1;attr.uid=0;attr.gid=0;attr.rdev=node.rdev;if(FS.isDir(node.mode)){attr.size=4096}else if(FS.isFile(node.mode)){attr.size=node.usedBytes}else if(FS.isLink(node.mode)){attr.size=node.link.length}else{attr.size=0}attr.atime=new Date(node.atime);attr.mtime=new Date(node.mtime);attr.ctime=new Date(node.ctime);attr.blksize=4096;attr.blocks=Math.ceil(attr.size/attr.blksize);return attr},setattr(node,attr){for(const key of["mode","atime","mtime","ctime"]){if(attr[key]!=null){node[key]=attr[key]}}if(attr.size!==undefined){MEMFS.resizeFileStorage(node,attr.size)}},lookup(parent,name){if(!MEMFS.doesNotExistError){MEMFS.doesNotExistError=new FS.ErrnoError(44);MEMFS.doesNotExistError.stack="<generic error, no stack>"}throw MEMFS.doesNotExistError},mknod(parent,name,mode,dev){return MEMFS.createNode(parent,name,mode,dev)},rename(old_node,new_dir,new_name){var new_node;try{new_node=FS.lookupNode(new_dir,new_name)}catch(e){}if(new_node){if(FS.isDir(old_node.mode)){for(var i in new_node.contents){throw new FS.ErrnoError(55)}}FS.hashRemoveNode(new_node)}delete old_node.parent.contents[old_node.name];new_dir.contents[new_name]=old_node;old_node.name=new_name;new_dir.ctime=new_dir.mtime=old_node.parent.ctime=old_node.parent.mtime=Date.now()},unlink(parent,name){delete parent.contents[name];parent.ctime=parent.mtime=Date.now()},rmdir(parent,name){var node=FS.lookupNode(parent,name);for(var i in node.contents){throw new FS.ErrnoError(55)}delete parent.contents[name];parent.ctime=parent.mtime=Date.now()},readdir(node){return[".","..",...Object.keys(node.contents)]},symlink(parent,newname,oldpath){var node=MEMFS.createNode(parent,newname,511|40960,0);node.link=oldpath;return node},readlink(node){if(!FS.isLink(node.mode)){throw new FS.ErrnoError(28)}return node.link}},stream_ops:{read(stream,buffer,offset,length,position){var contents=stream.node.contents;if(position>=stream.node.usedBytes)return 0;var size=Math.min(stream.node.usedBytes-position,length);if(size>8&&contents.subarray){buffer.set(contents.subarray(position,position+size),offset)}else{for(var i=0;i<size;i++)buffer[offset+i]=contents[position+i]}return size},write(stream,buffer,offset,length,position,canOwn){if(buffer.buffer===(growMemViews(),HEAP8).buffer){canOwn=false}if(!length)return 0;var node=stream.node;node.mtime=node.ctime=Date.now();if(buffer.subarray&&(!node.contents||node.contents.subarray)){if(canOwn){node.contents=buffer.subarray(offset,offset+length);node.usedBytes=length;return length}else if(node.usedBytes===0&&position===0){node.contents=buffer.slice(offset,offset+length);node.usedBytes=length;return length}else if(position+length<=node.usedBytes){node.contents.set(buffer.subarray(offset,offset+length),position);return length}}MEMFS.expandFileStorage(node,position+length);if(node.contents.subarray&&buffer.subarray){node.contents.set(buffer.subarray(offset,offset+length),position)}else{for(var i=0;i<length;i++){node.contents[position+i]=buffer[offset+i]}}node.usedBytes=Math.max(node.usedBytes,position+length);return length},llseek(stream,offset,whence){var position=offset;if(whence===1){position+=stream.position}else if(whence===2){if(FS.isFile(stream.node.mode)){position+=stream.node.usedBytes}}if(position<0){throw new FS.ErrnoError(28)}return position},mmap(stream,length,position,prot,flags){if(!FS.isFile(stream.node.mode)){throw new FS.ErrnoError(43)}var ptr;var allocated;var contents=stream.node.contents;if(!(flags&2)&&contents&&contents.buffer===(growMemViews(),HEAP8).buffer){allocated=false;ptr=contents.byteOffset}else{allocated=true;ptr=mmapAlloc(length);if(!ptr){throw new FS.ErrnoError(48)}if(contents){if(position>0||position+length<contents.length){if(contents.subarray){contents=contents.subarray(position,position+length)}else{contents=Array.prototype.slice.call(contents,position,position+length)}}(growMemViews(),HEAP8).set(contents,ptr)}}return{ptr,allocated}},msync(stream,buffer,offset,length,mmapFlags){MEMFS.stream_ops.write(stream,buffer,0,length,offset,false);return 0}}};var FS_modeStringToFlags=str=>{var flagModes={r:0,"r+":2,w:512|64|1,"w+":512|64|2,a:1024|64|1,"a+":1024|64|2};var flags=flagModes[str];if(typeof flags=="undefined"){throw new Error(`Unknown file open mode: ${str}`)}return flags};var FS_getMode=(canRead,canWrite)=>{var mode=0;if(canRead)mode|=292|73;if(canWrite)mode|=146;return mode};var asyncLoad=async url=>{var arrayBuffer=await readAsync(url);return new Uint8Array(arrayBuffer)};var FS_createDataFile=(...args)=>FS.createDataFile(...args);var getUniqueRunDependency=id=>id;var preloadPlugins=[];var FS_handledByPreloadPlugin=async(byteArray,fullname)=>{if(typeof Browser!="undefined")Browser.init();for(var plugin of preloadPlugins){if(plugin["canHandle"](fullname)){return plugin["handle"](byteArray,fullname)}}return byteArray};var FS_preloadFile=async(parent,name,url,canRead,canWrite,dontCreateFile,canOwn,preFinish)=>{var fullname=name?PATH_FS.resolve(PATH.join2(parent,name)):parent;var dep=getUniqueRunDependency(`cp ${fullname}`);addRunDependency(dep);try{var byteArray=url;if(typeof url=="string"){byteArray=await asyncLoad(url)}byteArray=await FS_handledByPreloadPlugin(byteArray,fullname);preFinish?.();if(!dontCreateFile){FS_createDataFile(parent,name,byteArray,canRead,canWrite,canOwn)}}finally{removeRunDependency(dep)}};var FS_createPreloadedFile=(parent,name,url,canRead,canWrite,onload,onerror,dontCreateFile,canOwn,preFinish)=>{FS_preloadFile(parent,name,url,canRead,canWrite,dontCreateFile,canOwn,preFinish).then(onload).catch(onerror)};var FS={root:null,mounts:[],devices:{},streams:[],nextInode:1,nameTable:null,currentPath:"/",initialized:false,ignorePermissions:true,filesystems:null,syncFSRequests:0,readFiles:{},ErrnoError:class{name="ErrnoError";constructor(errno){this.errno=errno}},FSStream:class{shared={};get object(){return this.node}set object(val){this.node=val}get isRead(){return(this.flags&2097155)!==1}get isWrite(){return(this.flags&2097155)!==0}get isAppend(){return this.flags&1024}get flags(){return this.shared.flags}set flags(val){this.shared.flags=val}get position(){return this.shared.position}set position(val){this.shared.position=val}},FSNode:class{node_ops={};stream_ops={};readMode=292|73;writeMode=146;mounted=null;constructor(parent,name,mode,rdev){if(!parent){parent=this}this.parent=parent;this.mount=parent.mount;this.id=FS.nextInode++;this.name=name;this.mode=mode;this.rdev=rdev;this.atime=this.mtime=this.ctime=Date.now()}get read(){return(this.mode&this.readMode)===this.readMode}set read(val){val?this.mode|=this.readMode:this.mode&=~this.readMode}get write(){return(this.mode&this.writeMode)===this.writeMode}set write(val){val?this.mode|=this.writeMode:this.mode&=~this.writeMode}get isFolder(){return FS.isDir(this.mode)}get isDevice(){return FS.isChrdev(this.mode)}},lookupPath(path,opts={}){if(!path){throw new FS.ErrnoError(44)}opts.follow_mount??=true;if(!PATH.isAbs(path)){path=FS.cwd()+"/"+path}linkloop:for(var nlinks=0;nlinks<40;nlinks++){var parts=path.split("/").filter(p=>!!p);var current=FS.root;var current_path="/";for(var i=0;i<parts.length;i++){var islast=i===parts.length-1;if(islast&&opts.parent){break}if(parts[i]==="."){continue}if(parts[i]===".."){current_path=PATH.dirname(current_path);if(FS.isRoot(current)){path=current_path+"/"+parts.slice(i+1).join("/");nlinks--;continue linkloop}else{current=current.parent}continue}current_path=PATH.join2(current_path,parts[i]);try{current=FS.lookupNode(current,parts[i])}catch(e){if(e?.errno===44&&islast&&opts.noent_okay){return{path:current_path}}throw e}if(FS.isMountpoint(current)&&(!islast||opts.follow_mount)){current=current.mounted.root}if(FS.isLink(current.mode)&&(!islast||opts.follow)){if(!current.node_ops.readlink){throw new FS.ErrnoError(52)}var link=current.node_ops.readlink(current);if(!PATH.isAbs(link)){link=PATH.dirname(current_path)+"/"+link}path=link+"/"+parts.slice(i+1).join("/");continue linkloop}}return{path:current_path,node:current}}throw new FS.ErrnoError(32)},getPath(node){var path;while(true){if(FS.isRoot(node)){var mount=node.mount.mountpoint;if(!path)return mount;return mount[mount.length-1]!=="/"?`${mount}/${path}`:mount+path}path=path?`${node.name}/${path}`:node.name;node=node.parent}},hashName(parentid,name){var hash=0;for(var i=0;i<name.length;i++){hash=(hash<<5)-hash+name.charCodeAt(i)|0}return(parentid+hash>>>0)%FS.nameTable.length},hashAddNode(node){var hash=FS.hashName(node.parent.id,node.name);node.name_next=FS.nameTable[hash];FS.nameTable[hash]=node},hashRemoveNode(node){var hash=FS.hashName(node.parent.id,node.name);if(FS.nameTable[hash]===node){FS.nameTable[hash]=node.name_next}else{var current=FS.nameTable[hash];while(current){if(current.name_next===node){current.name_next=node.name_next;break}current=current.name_next}}},lookupNode(parent,name){var errCode=FS.mayLookup(parent);if(errCode){throw new FS.ErrnoError(errCode)}var hash=FS.hashName(parent.id,name);for(var node=FS.nameTable[hash];node;node=node.name_next){var nodeName=node.name;if(node.parent.id===parent.id&&nodeName===name){return node}}return FS.lookup(parent,name)},createNode(parent,name,mode,rdev){var node=new FS.FSNode(parent,name,mode,rdev);FS.hashAddNode(node);return node},destroyNode(node){FS.hashRemoveNode(node)},isRoot(node){return node===node.parent},isMountpoint(node){return!!node.mounted},isFile(mode){return(mode&61440)===32768},isDir(mode){return(mode&61440)===16384},isLink(mode){return(mode&61440)===40960},isChrdev(mode){return(mode&61440)===8192},isBlkdev(mode){return(mode&61440)===24576},isFIFO(mode){return(mode&61440)===4096},isSocket(mode){return(mode&49152)===49152},flagsToPermissionString(flag){var perms=["r","w","rw"][flag&3];if(flag&512){perms+="w"}return perms},nodePermissions(node,perms){if(FS.ignorePermissions){return 0}if(perms.includes("r")&&!(node.mode&292)){return 2}else if(perms.includes("w")&&!(node.mode&146)){return 2}else if(perms.includes("x")&&!(node.mode&73)){return 2}return 0},mayLookup(dir){if(!FS.isDir(dir.mode))return 54;var errCode=FS.nodePermissions(dir,"x");if(errCode)return errCode;if(!dir.node_ops.lookup)return 2;return 0},mayCreate(dir,name){if(!FS.isDir(dir.mode)){return 54}try{var node=FS.lookupNode(dir,name);return 20}catch(e){}return FS.nodePermissions(dir,"wx")},mayDelete(dir,name,isdir){var node;try{node=FS.lookupNode(dir,name)}catch(e){return e.errno}var errCode=FS.nodePermissions(dir,"wx");if(errCode){return errCode}if(isdir){if(!FS.isDir(node.mode)){return 54}if(FS.isRoot(node)||FS.getPath(node)===FS.cwd()){return 10}}else{if(FS.isDir(node.mode)){return 31}}return 0},mayOpen(node,flags){if(!node){return 44}if(FS.isLink(node.mode)){return 32}else if(FS.isDir(node.mode)){if(FS.flagsToPermissionString(flags)!=="r"||flags&(512|64)){return 31}}return FS.nodePermissions(node,FS.flagsToPermissionString(flags))},checkOpExists(op,err){if(!op){throw new FS.ErrnoError(err)}return op},MAX_OPEN_FDS:4096,nextfd(){for(var fd=0;fd<=FS.MAX_OPEN_FDS;fd++){if(!FS.streams[fd]){return fd}}throw new FS.ErrnoError(33)},getStreamChecked(fd){var stream=FS.getStream(fd);if(!stream){throw new FS.ErrnoError(8)}return stream},getStream:fd=>FS.streams[fd],createStream(stream,fd=-1){stream=Object.assign(new FS.FSStream,stream);if(fd==-1){fd=FS.nextfd()}stream.fd=fd;FS.streams[fd]=stream;return stream},closeStream(fd){FS.streams[fd]=null},dupStream(origStream,fd=-1){var stream=FS.createStream(origStream,fd);stream.stream_ops?.dup?.(stream);return stream},doSetAttr(stream,node,attr){var setattr=stream?.stream_ops.setattr;var arg=setattr?stream:node;setattr??=node.node_ops.setattr;FS.checkOpExists(setattr,63);setattr(arg,attr)},chrdev_stream_ops:{open(stream){var device=FS.getDevice(stream.node.rdev);stream.stream_ops=device.stream_ops;stream.stream_ops.open?.(stream)},llseek(){throw new FS.ErrnoError(70)}},major:dev=>dev>>8,minor:dev=>dev&255,makedev:(ma,mi)=>ma<<8|mi,registerDevice(dev,ops){FS.devices[dev]={stream_ops:ops}},getDevice:dev=>FS.devices[dev],getMounts(mount){var mounts=[];var check=[mount];while(check.length){var m=check.pop();mounts.push(m);check.push(...m.mounts)}return mounts},syncfs(populate,callback){if(typeof populate=="function"){callback=populate;populate=false}FS.syncFSRequests++;if(FS.syncFSRequests>1){err(`warning: ${FS.syncFSRequests} FS.syncfs operations in flight at once, probably just doing extra work`)}var mounts=FS.getMounts(FS.root.mount);var completed=0;function doCallback(errCode){FS.syncFSRequests--;return callback(errCode)}function done(errCode){if(errCode){if(!done.errored){done.errored=true;return doCallback(errCode)}return}if(++completed>=mounts.length){doCallback(null)}}for(var mount of mounts){if(mount.type.syncfs){mount.type.syncfs(mount,populate,done)}else{done(null)}}},mount(type,opts,mountpoint){var root=mountpoint==="/";var pseudo=!mountpoint;var node;if(root&&FS.root){throw new FS.ErrnoError(10)}else if(!root&&!pseudo){var lookup=FS.lookupPath(mountpoint,{follow_mount:false});mountpoint=lookup.path;node=lookup.node;if(FS.isMountpoint(node)){throw new FS.ErrnoError(10)}if(!FS.isDir(node.mode)){throw new FS.ErrnoError(54)}}var mount={type,opts,mountpoint,mounts:[]};var mountRoot=type.mount(mount);mountRoot.mount=mount;mount.root=mountRoot;if(root){FS.root=mountRoot}else if(node){node.mounted=mount;if(node.mount){node.mount.mounts.push(mount)}}return mountRoot},unmount(mountpoint){var lookup=FS.lookupPath(mountpoint,{follow_mount:false});if(!FS.isMountpoint(lookup.node)){throw new FS.ErrnoError(28)}var node=lookup.node;var mount=node.mounted;var mounts=FS.getMounts(mount);for(var[hash,current]of Object.entries(FS.nameTable)){while(current){var next=current.name_next;if(mounts.includes(current.mount)){FS.destroyNode(current)}current=next}}node.mounted=null;var idx=node.mount.mounts.indexOf(mount);node.mount.mounts.splice(idx,1)},lookup(parent,name){return parent.node_ops.lookup(parent,name)},mknod(path,mode,dev){var lookup=FS.lookupPath(path,{parent:true});var parent=lookup.node;var name=PATH.basename(path);if(!name){throw new FS.ErrnoError(28)}if(name==="."||name===".."){throw new FS.ErrnoError(20)}var errCode=FS.mayCreate(parent,name);if(errCode){throw new FS.ErrnoError(errCode)}if(!parent.node_ops.mknod){throw new FS.ErrnoError(63)}return parent.node_ops.mknod(parent,name,mode,dev)},statfs(path){return FS.statfsNode(FS.lookupPath(path,{follow:true}).node)},statfsStream(stream){return FS.statfsNode(stream.node)},statfsNode(node){var rtn={bsize:4096,frsize:4096,blocks:1e6,bfree:5e5,bavail:5e5,files:FS.nextInode,ffree:FS.nextInode-1,fsid:42,flags:2,namelen:255};if(node.node_ops.statfs){Object.assign(rtn,node.node_ops.statfs(node.mount.opts.root))}return rtn},create(path,mode=438){mode&=4095;mode|=32768;return FS.mknod(path,mode,0)},mkdir(path,mode=511){mode&=511|512;mode|=16384;return FS.mknod(path,mode,0)},mkdirTree(path,mode){var dirs=path.split("/");var d="";for(var dir of dirs){if(!dir)continue;if(d||PATH.isAbs(path))d+="/";d+=dir;try{FS.mkdir(d,mode)}catch(e){if(e.errno!=20)throw e}}},mkdev(path,mode,dev){if(typeof dev=="undefined"){dev=mode;mode=438}mode|=8192;return FS.mknod(path,mode,dev)},symlink(oldpath,newpath){if(!PATH_FS.resolve(oldpath)){throw new FS.ErrnoError(44)}var lookup=FS.lookupPath(newpath,{parent:true});var parent=lookup.node;if(!parent){throw new FS.ErrnoError(44)}var newname=PATH.basename(newpath);var errCode=FS.mayCreate(parent,newname);if(errCode){throw new FS.ErrnoError(errCode)}if(!parent.node_ops.symlink){throw new FS.ErrnoError(63)}return parent.node_ops.symlink(parent,newname,oldpath)},rename(old_path,new_path){var old_dirname=PATH.dirname(old_path);var new_dirname=PATH.dirname(new_path);var old_name=PATH.basename(old_path);var new_name=PATH.basename(new_path);var lookup,old_dir,new_dir;lookup=FS.lookupPath(old_path,{parent:true});old_dir=lookup.node;lookup=FS.lookupPath(new_path,{parent:true});new_dir=lookup.node;if(!old_dir||!new_dir)throw new FS.ErrnoError(44);if(old_dir.mount!==new_dir.mount){throw new FS.ErrnoError(75)}var old_node=FS.lookupNode(old_dir,old_name);var relative=PATH_FS.relative(old_path,new_dirname);if(relative.charAt(0)!=="."){throw new FS.ErrnoError(28)}relative=PATH_FS.relative(new_path,old_dirname);if(relative.charAt(0)!=="."){throw new FS.ErrnoError(55)}var new_node;try{new_node=FS.lookupNode(new_dir,new_name)}catch(e){}if(old_node===new_node){return}var isdir=FS.isDir(old_node.mode);var errCode=FS.mayDelete(old_dir,old_name,isdir);if(errCode){throw new FS.ErrnoError(errCode)}errCode=new_node?FS.mayDelete(new_dir,new_name,isdir):FS.mayCreate(new_dir,new_name);if(errCode){throw new FS.ErrnoError(errCode)}if(!old_dir.node_ops.rename){throw new FS.ErrnoError(63)}if(FS.isMountpoint(old_node)||new_node&&FS.isMountpoint(new_node)){throw new FS.ErrnoError(10)}if(new_dir!==old_dir){errCode=FS.nodePermissions(old_dir,"w");if(errCode){throw new FS.ErrnoError(errCode)}}FS.hashRemoveNode(old_node);try{old_dir.node_ops.rename(old_node,new_dir,new_name);old_node.parent=new_dir}catch(e){throw e}finally{FS.hashAddNode(old_node)}},rmdir(path){var lookup=FS.lookupPath(path,{parent:true});var parent=lookup.node;var name=PATH.basename(path);var node=FS.lookupNode(parent,name);var errCode=FS.mayDelete(parent,name,true);if(errCode){throw new FS.ErrnoError(errCode)}if(!parent.node_ops.rmdir){throw new FS.ErrnoError(63)}if(FS.isMountpoint(node)){throw new FS.ErrnoError(10)}parent.node_ops.rmdir(parent,name);FS.destroyNode(node)},readdir(path){var lookup=FS.lookupPath(path,{follow:true});var node=lookup.node;var readdir=FS.checkOpExists(node.node_ops.readdir,54);return readdir(node)},unlink(path){var lookup=FS.lookupPath(path,{parent:true});var parent=lookup.node;if(!parent){throw new FS.ErrnoError(44)}var name=PATH.basename(path);var node=FS.lookupNode(parent,name);var errCode=FS.mayDelete(parent,name,false);if(errCode){throw new FS.ErrnoError(errCode)}if(!parent.node_ops.unlink){throw new FS.ErrnoError(63)}if(FS.isMountpoint(node)){throw new FS.ErrnoError(10)}parent.node_ops.unlink(parent,name);FS.destroyNode(node)},readlink(path){var lookup=FS.lookupPath(path);var link=lookup.node;if(!link){throw new FS.ErrnoError(44)}if(!link.node_ops.readlink){throw new FS.ErrnoError(28)}return link.node_ops.readlink(link)},stat(path,dontFollow){var lookup=FS.lookupPath(path,{follow:!dontFollow});var node=lookup.node;var getattr=FS.checkOpExists(node.node_ops.getattr,63);return getattr(node)},fstat(fd){var stream=FS.getStreamChecked(fd);var node=stream.node;var getattr=stream.stream_ops.getattr;var arg=getattr?stream:node;getattr??=node.node_ops.getattr;FS.checkOpExists(getattr,63);return getattr(arg)},lstat(path){return FS.stat(path,true)},doChmod(stream,node,mode,dontFollow){FS.doSetAttr(stream,node,{mode:mode&4095|node.mode&~4095,ctime:Date.now(),dontFollow})},chmod(path,mode,dontFollow){var node;if(typeof path=="string"){var lookup=FS.lookupPath(path,{follow:!dontFollow});node=lookup.node}else{node=path}FS.doChmod(null,node,mode,dontFollow)},lchmod(path,mode){FS.chmod(path,mode,true)},fchmod(fd,mode){var stream=FS.getStreamChecked(fd);FS.doChmod(stream,stream.node,mode,false)},doChown(stream,node,dontFollow){FS.doSetAttr(stream,node,{timestamp:Date.now(),dontFollow})},chown(path,uid,gid,dontFollow){var node;if(typeof path=="string"){var lookup=FS.lookupPath(path,{follow:!dontFollow});node=lookup.node}else{node=path}FS.doChown(null,node,dontFollow)},lchown(path,uid,gid){FS.chown(path,uid,gid,true)},fchown(fd,uid,gid){var stream=FS.getStreamChecked(fd);FS.doChown(stream,stream.node,false)},doTruncate(stream,node,len){if(FS.isDir(node.mode)){throw new FS.ErrnoError(31)}if(!FS.isFile(node.mode)){throw new FS.ErrnoError(28)}var errCode=FS.nodePermissions(node,"w");if(errCode){throw new FS.ErrnoError(errCode)}FS.doSetAttr(stream,node,{size:len,timestamp:Date.now()})},truncate(path,len){if(len<0){throw new FS.ErrnoError(28)}var node;if(typeof path=="string"){var lookup=FS.lookupPath(path,{follow:true});node=lookup.node}else{node=path}FS.doTruncate(null,node,len)},ftruncate(fd,len){var stream=FS.getStreamChecked(fd);if(len<0||(stream.flags&2097155)===0){throw new FS.ErrnoError(28)}FS.doTruncate(stream,stream.node,len)},utime(path,atime,mtime){var lookup=FS.lookupPath(path,{follow:true});var node=lookup.node;var setattr=FS.checkOpExists(node.node_ops.setattr,63);setattr(node,{atime,mtime})},open(path,flags,mode=438){if(path===""){throw new FS.ErrnoError(44)}flags=typeof flags=="string"?FS_modeStringToFlags(flags):flags;if(flags&64){mode=mode&4095|32768}else{mode=0}var node;var isDirPath;if(typeof path=="object"){node=path}else{isDirPath=path.endsWith("/");var lookup=FS.lookupPath(path,{follow:!(flags&131072),noent_okay:true});node=lookup.node;path=lookup.path}var created=false;if(flags&64){if(node){if(flags&128){throw new FS.ErrnoError(20)}}else if(isDirPath){throw new FS.ErrnoError(31)}else{node=FS.mknod(path,mode|511,0);created=true}}if(!node){throw new FS.ErrnoError(44)}if(FS.isChrdev(node.mode)){flags&=~512}if(flags&65536&&!FS.isDir(node.mode)){throw new FS.ErrnoError(54)}if(!created){var errCode=FS.mayOpen(node,flags);if(errCode){throw new FS.ErrnoError(errCode)}}if(flags&512&&!created){FS.truncate(node,0)}flags&=~(128|512|131072);var stream=FS.createStream({node,path:FS.getPath(node),flags,seekable:true,position:0,stream_ops:node.stream_ops,ungotten:[],error:false});if(stream.stream_ops.open){stream.stream_ops.open(stream)}if(created){FS.chmod(node,mode&511)}if(Module["logReadFiles"]&&!(flags&1)){if(!(path in FS.readFiles)){FS.readFiles[path]=1}}return stream},close(stream){if(FS.isClosed(stream)){throw new FS.ErrnoError(8)}if(stream.getdents)stream.getdents=null;try{if(stream.stream_ops.close){stream.stream_ops.close(stream)}}catch(e){throw e}finally{FS.closeStream(stream.fd)}stream.fd=null},isClosed(stream){return stream.fd===null},llseek(stream,offset,whence){if(FS.isClosed(stream)){throw new FS.ErrnoError(8)}if(!stream.seekable||!stream.stream_ops.llseek){throw new FS.ErrnoError(70)}if(whence!=0&&whence!=1&&whence!=2){throw new FS.ErrnoError(28)}stream.position=stream.stream_ops.llseek(stream,offset,whence);stream.ungotten=[];return stream.position},read(stream,buffer,offset,length,position){if(length<0||position<0){throw new FS.ErrnoError(28)}if(FS.isClosed(stream)){throw new FS.ErrnoError(8)}if((stream.flags&2097155)===1){throw new FS.ErrnoError(8)}if(FS.isDir(stream.node.mode)){throw new FS.ErrnoError(31)}if(!stream.stream_ops.read){throw new FS.ErrnoError(28)}var seeking=typeof position!="undefined";if(!seeking){position=stream.position}else if(!stream.seekable){throw new FS.ErrnoError(70)}var bytesRead=stream.stream_ops.read(stream,buffer,offset,length,position);if(!seeking)stream.position+=bytesRead;return bytesRead},write(stream,buffer,offset,length,position,canOwn){if(length<0||position<0){throw new FS.ErrnoError(28)}if(FS.isClosed(stream)){throw new FS.ErrnoError(8)}if((stream.flags&2097155)===0){throw new FS.ErrnoError(8)}if(FS.isDir(stream.node.mode)){throw new FS.ErrnoError(31)}if(!stream.stream_ops.write){throw new FS.ErrnoError(28)}if(stream.seekable&&stream.flags&1024){FS.llseek(stream,0,2)}var seeking=typeof position!="undefined";if(!seeking){position=stream.position}else if(!stream.seekable){throw new FS.ErrnoError(70)}var bytesWritten=stream.stream_ops.write(stream,buffer,offset,length,position,canOwn);if(!seeking)stream.position+=bytesWritten;return bytesWritten},mmap(stream,length,position,prot,flags){if((prot&2)!==0&&(flags&2)===0&&(stream.flags&2097155)!==2){throw new FS.ErrnoError(2)}if((stream.flags&2097155)===1){throw new FS.ErrnoError(2)}if(!stream.stream_ops.mmap){throw new FS.ErrnoError(43)}if(!length){throw new FS.ErrnoError(28)}return stream.stream_ops.mmap(stream,length,position,prot,flags)},msync(stream,buffer,offset,length,mmapFlags){if(!stream.stream_ops.msync){return 0}return stream.stream_ops.msync(stream,buffer,offset,length,mmapFlags)},ioctl(stream,cmd,arg){if(!stream.stream_ops.ioctl){throw new FS.ErrnoError(59)}return stream.stream_ops.ioctl(stream,cmd,arg)},readFile(path,opts={}){opts.flags=opts.flags||0;opts.encoding=opts.encoding||"binary";if(opts.encoding!=="utf8"&&opts.encoding!=="binary"){abort(`Invalid encoding type "${opts.encoding}"`)}var stream=FS.open(path,opts.flags);var stat=FS.stat(path);var length=stat.size;var buf=new Uint8Array(length);FS.read(stream,buf,0,length,0);if(opts.encoding==="utf8"){buf=UTF8ArrayToString(buf)}FS.close(stream);return buf},writeFile(path,data,opts={}){opts.flags=opts.flags||577;var stream=FS.open(path,opts.flags,opts.mode);if(typeof data=="string"){data=new Uint8Array(intArrayFromString(data,true))}if(ArrayBuffer.isView(data)){FS.write(stream,data,0,data.byteLength,undefined,opts.canOwn)}else{abort("Unsupported data type")}FS.close(stream)},cwd:()=>FS.currentPath,chdir(path){var lookup=FS.lookupPath(path,{follow:true});if(lookup.node===null){throw new FS.ErrnoError(44)}if(!FS.isDir(lookup.node.mode)){throw new FS.ErrnoError(54)}var errCode=FS.nodePermissions(lookup.node,"x");if(errCode){throw new FS.ErrnoError(errCode)}FS.currentPath=lookup.path},createDefaultDirectories(){FS.mkdir("/tmp");FS.mkdir("/home");FS.mkdir("/home/web_user")},createDefaultDevices(){FS.mkdir("/dev");FS.registerDevice(FS.makedev(1,3),{read:()=>0,write:(stream,buffer,offset,length,pos)=>length,llseek:()=>0});FS.mkdev("/dev/null",FS.makedev(1,3));TTY.register(FS.makedev(5,0),TTY.default_tty_ops);TTY.register(FS.makedev(6,0),TTY.default_tty1_ops);FS.mkdev("/dev/tty",FS.makedev(5,0));FS.mkdev("/dev/tty1",FS.makedev(6,0));var randomBuffer=new Uint8Array(1024),randomLeft=0;var randomByte=()=>{if(randomLeft===0){randomFill(randomBuffer);randomLeft=randomBuffer.byteLength}return randomBuffer[--randomLeft]};FS.createDevice("/dev","random",randomByte);FS.createDevice("/dev","urandom",randomByte);FS.mkdir("/dev/shm");FS.mkdir("/dev/shm/tmp")},createSpecialDirectories(){FS.mkdir("/proc");var proc_self=FS.mkdir("/proc/self");FS.mkdir("/proc/self/fd");FS.mount({mount(){var node=FS.createNode(proc_self,"fd",16895,73);node.stream_ops={llseek:MEMFS.stream_ops.llseek};node.node_ops={lookup(parent,name){var fd=+name;var stream=FS.getStreamChecked(fd);var ret={parent:null,mount:{mountpoint:"fake"},node_ops:{readlink:()=>stream.path},id:fd+1};ret.parent=ret;return ret},readdir(){return Array.from(FS.streams.entries()).filter(([k,v])=>v).map(([k,v])=>k.toString())}};return node}},{},"/proc/self/fd")},createStandardStreams(input,output,error){if(input){FS.createDevice("/dev","stdin",input)}else{FS.symlink("/dev/tty","/dev/stdin")}if(output){FS.createDevice("/dev","stdout",null,output)}else{FS.symlink("/dev/tty","/dev/stdout")}if(error){FS.createDevice("/dev","stderr",null,error)}else{FS.symlink("/dev/tty1","/dev/stderr")}var stdin=FS.open("/dev/stdin",0);var stdout=FS.open("/dev/stdout",1);var stderr=FS.open("/dev/stderr",1)},staticInit(){FS.nameTable=new Array(4096);FS.mount(MEMFS,{},"/");FS.createDefaultDirectories();FS.createDefaultDevices();FS.createSpecialDirectories();FS.filesystems={MEMFS}},init(input,output,error){FS.initialized=true;input??=Module["stdin"];output??=Module["stdout"];error??=Module["stderr"];FS.createStandardStreams(input,output,error)},quit(){FS.initialized=false;for(var stream of FS.streams){if(stream){FS.close(stream)}}},findObject(path,dontResolveLastLink){var ret=FS.analyzePath(path,dontResolveLastLink);if(!ret.exists){return null}return ret.object},analyzePath(path,dontResolveLastLink){try{var lookup=FS.lookupPath(path,{follow:!dontResolveLastLink});path=lookup.path}catch(e){}var ret={isRoot:false,exists:false,error:0,name:null,path:null,object:null,parentExists:false,parentPath:null,parentObject:null};try{var lookup=FS.lookupPath(path,{parent:true});ret.parentExists=true;ret.parentPath=lookup.path;ret.parentObject=lookup.node;ret.name=PATH.basename(path);lookup=FS.lookupPath(path,{follow:!dontResolveLastLink});ret.exists=true;ret.path=lookup.path;ret.object=lookup.node;ret.name=lookup.node.name;ret.isRoot=lookup.path==="/"}catch(e){ret.error=e.errno}return ret},createPath(parent,path,canRead,canWrite){parent=typeof parent=="string"?parent:FS.getPath(parent);var parts=path.split("/").reverse();while(parts.length){var part=parts.pop();if(!part)continue;var current=PATH.join2(parent,part);try{FS.mkdir(current)}catch(e){if(e.errno!=20)throw e}parent=current}return current},createFile(parent,name,properties,canRead,canWrite){var path=PATH.join2(typeof parent=="string"?parent:FS.getPath(parent),name);var mode=FS_getMode(canRead,canWrite);return FS.create(path,mode)},createDataFile(parent,name,data,canRead,canWrite,canOwn){var path=name;if(parent){parent=typeof parent=="string"?parent:FS.getPath(parent);path=name?PATH.join2(parent,name):parent}var mode=FS_getMode(canRead,canWrite);var node=FS.create(path,mode);if(data){if(typeof data=="string"){var arr=new Array(data.length);for(var i=0,len=data.length;i<len;++i)arr[i]=data.charCodeAt(i);data=arr}FS.chmod(node,mode|146);var stream=FS.open(node,577);FS.write(stream,data,0,data.length,0,canOwn);FS.close(stream);FS.chmod(node,mode)}},createDevice(parent,name,input,output){var path=PATH.join2(typeof parent=="string"?parent:FS.getPath(parent),name);var mode=FS_getMode(!!input,!!output);FS.createDevice.major??=64;var dev=FS.makedev(FS.createDevice.major++,0);FS.registerDevice(dev,{open(stream){stream.seekable=false},close(stream){if(output?.buffer?.length){output(10)}},read(stream,buffer,offset,length,pos){var bytesRead=0;for(var i=0;i<length;i++){var result;try{result=input()}catch(e){throw new FS.ErrnoError(29)}if(result===undefined&&bytesRead===0){throw new FS.ErrnoError(6)}if(result===null||result===undefined)break;bytesRead++;buffer[offset+i]=result}if(bytesRead){stream.node.atime=Date.now()}return bytesRead},write(stream,buffer,offset,length,pos){for(var i=0;i<length;i++){try{output(buffer[offset+i])}catch(e){throw new FS.ErrnoError(29)}}if(length){stream.node.mtime=stream.node.ctime=Date.now()}return i}});return FS.mkdev(path,mode,dev)},forceLoadFile(obj){if(obj.isDevice||obj.isFolder||obj.link||obj.contents)return true;if(globalThis.XMLHttpRequest){abort("Lazy loading should have been performed (contents set) in createLazyFile, but it was not. Lazy loading only works in web workers. Use --embed-file or --preload-file in emcc on the main thread.")}else{try{obj.contents=readBinary(obj.url)}catch(e){throw new FS.ErrnoError(29)}}},createLazyFile(parent,name,url,canRead,canWrite){class LazyUint8Array{lengthKnown=false;chunks=[];get(idx){if(idx>this.length-1||idx<0){return undefined}var chunkOffset=idx%this.chunkSize;var chunkNum=idx/this.chunkSize|0;return this.getter(chunkNum)[chunkOffset]}setDataGetter(getter){this.getter=getter}cacheLength(){var xhr=new XMLHttpRequest;xhr.open("HEAD",url,false);xhr.send(null);if(!(xhr.status>=200&&xhr.status<300||xhr.status===304))abort("Couldn\'t load "+url+". Status: "+xhr.status);var datalength=Number(xhr.getResponseHeader("Content-length"));var header;var hasByteServing=(header=xhr.getResponseHeader("Accept-Ranges"))&&header==="bytes";var usesGzip=(header=xhr.getResponseHeader("Content-Encoding"))&&header==="gzip";var chunkSize=1024*1024;if(!hasByteServing)chunkSize=datalength;var doXHR=(from,to)=>{if(from>to)abort("invalid range ("+from+", "+to+") or no bytes requested!");if(to>datalength-1)abort("only "+datalength+" bytes available! programmer error!");var xhr=new XMLHttpRequest;xhr.open("GET",url,false);if(datalength!==chunkSize)xhr.setRequestHeader("Range","bytes="+from+"-"+to);xhr.responseType="arraybuffer";if(xhr.overrideMimeType){xhr.overrideMimeType("text/plain; charset=x-user-defined")}xhr.send(null);if(!(xhr.status>=200&&xhr.status<300||xhr.status===304))abort("Couldn\'t load "+url+". Status: "+xhr.status);if(xhr.response!==undefined){return new Uint8Array(xhr.response||[])}return intArrayFromString(xhr.responseText||"",true)};var lazyArray=this;lazyArray.setDataGetter(chunkNum=>{var start=chunkNum*chunkSize;var end=(chunkNum+1)*chunkSize-1;end=Math.min(end,datalength-1);if(typeof lazyArray.chunks[chunkNum]=="undefined"){lazyArray.chunks[chunkNum]=doXHR(start,end)}if(typeof lazyArray.chunks[chunkNum]=="undefined")abort("doXHR failed!");return lazyArray.chunks[chunkNum]});if(usesGzip||!datalength){chunkSize=datalength=1;datalength=this.getter(0).length;chunkSize=datalength;out("LazyFiles on gzip forces download of the whole file when length is accessed")}this._length=datalength;this._chunkSize=chunkSize;this.lengthKnown=true}get length(){if(!this.lengthKnown){this.cacheLength()}return this._length}get chunkSize(){if(!this.lengthKnown){this.cacheLength()}return this._chunkSize}}if(globalThis.XMLHttpRequest){if(!ENVIRONMENT_IS_WORKER)abort("Cannot do synchronous binary XHRs outside webworkers in modern browsers. Use --embed-file or --preload-file in emcc");var lazyArray=new LazyUint8Array;var properties={isDevice:false,contents:lazyArray}}else{var properties={isDevice:false,url}}var node=FS.createFile(parent,name,properties,canRead,canWrite);if(properties.contents){node.contents=properties.contents}else if(properties.url){node.contents=null;node.url=properties.url}Object.defineProperties(node,{usedBytes:{get:function(){return this.contents.length}}});var stream_ops={};for(const[key,fn]of Object.entries(node.stream_ops)){stream_ops[key]=(...args)=>{FS.forceLoadFile(node);return fn(...args)}}function writeChunks(stream,buffer,offset,length,position){var contents=stream.node.contents;if(position>=contents.length)return 0;var size=Math.min(contents.length-position,length);if(contents.slice){for(var i=0;i<size;i++){buffer[offset+i]=contents[position+i]}}else{for(var i=0;i<size;i++){buffer[offset+i]=contents.get(position+i)}}return size}stream_ops.read=(stream,buffer,offset,length,position)=>{FS.forceLoadFile(node);return writeChunks(stream,buffer,offset,length,position)};stream_ops.mmap=(stream,length,position,prot,flags)=>{FS.forceLoadFile(node);var ptr=mmapAlloc(length);if(!ptr){throw new FS.ErrnoError(48)}writeChunks(stream,(growMemViews(),HEAP8),ptr,length,position);return{ptr,allocated:true}};node.stream_ops=stream_ops;return node}};var UTF8ToString=(ptr,maxBytesToRead,ignoreNul)=>ptr?UTF8ArrayToString((growMemViews(),HEAPU8),ptr,maxBytesToRead,ignoreNul):"";var SYSCALLS={DEFAULT_POLLMASK:5,calculateAt(dirfd,path,allowEmpty){if(PATH.isAbs(path)){return path}var dir;if(dirfd===-100){dir=FS.cwd()}else{var dirstream=SYSCALLS.getStreamFromFD(dirfd);dir=dirstream.path}if(path.length==0){if(!allowEmpty){throw new FS.ErrnoError(44)}return dir}return dir+"/"+path},writeStat(buf,stat){(growMemViews(),HEAPU32)[buf/4]=stat.dev;(growMemViews(),HEAPU32)[(buf+4)/4]=stat.mode;(growMemViews(),HEAPU64)[(buf+8)/8]=BigInt(stat.nlink);(growMemViews(),HEAPU32)[(buf+16)/4]=stat.uid;(growMemViews(),HEAPU32)[(buf+20)/4]=stat.gid;(growMemViews(),HEAPU32)[(buf+24)/4]=stat.rdev;(growMemViews(),HEAP64)[(buf+32)/8]=BigInt(stat.size);(growMemViews(),HEAP32)[(buf+40)/4]=4096;(growMemViews(),HEAP32)[(buf+44)/4]=stat.blocks;var atime=stat.atime.getTime();var mtime=stat.mtime.getTime();var ctime=stat.ctime.getTime();(growMemViews(),HEAP64)[(buf+48)/8]=BigInt(Math.floor(atime/1e3));(growMemViews(),HEAPU64)[(buf+56)/8]=BigInt(atime%1e3*1e3*1e3);(growMemViews(),HEAP64)[(buf+64)/8]=BigInt(Math.floor(mtime/1e3));(growMemViews(),HEAPU64)[(buf+72)/8]=BigInt(mtime%1e3*1e3*1e3);(growMemViews(),HEAP64)[(buf+80)/8]=BigInt(Math.floor(ctime/1e3));(growMemViews(),HEAPU64)[(buf+88)/8]=BigInt(ctime%1e3*1e3*1e3);(growMemViews(),HEAP64)[(buf+96)/8]=BigInt(stat.ino);return 0},writeStatFs(buf,stats){(growMemViews(),HEAPU32)[(buf+8)/4]=stats.bsize;(growMemViews(),HEAPU32)[(buf+72)/4]=stats.bsize;(growMemViews(),HEAP64)[(buf+16)/8]=BigInt(stats.blocks);(growMemViews(),HEAP64)[(buf+24)/8]=BigInt(stats.bfree);(growMemViews(),HEAP64)[(buf+32)/8]=BigInt(stats.bavail);(growMemViews(),HEAP64)[(buf+40)/8]=BigInt(stats.files);(growMemViews(),HEAP64)[(buf+48)/8]=BigInt(stats.ffree);(growMemViews(),HEAPU32)[(buf+56)/4]=stats.fsid;(growMemViews(),HEAPU32)[(buf+80)/4]=stats.flags;(growMemViews(),HEAPU32)[(buf+64)/4]=stats.namelen},doMsync(addr,stream,len,flags,offset){if(!FS.isFile(stream.node.mode)){throw new FS.ErrnoError(43)}if(flags&2){return 0}var buffer=(growMemViews(),HEAPU8).slice(addr,addr+len);FS.msync(stream,buffer,offset,len,flags)},getStreamFromFD(fd){var stream=FS.getStreamChecked(fd);return stream},varargs:undefined,getStr(ptr){var ret=UTF8ToString(ptr);return ret}};function ___syscall_fcntl64(fd,cmd,varargs){if(ENVIRONMENT_IS_PTHREAD)return proxyToMainThread(3,0,1,fd,cmd,varargs);varargs=bigintToI53Checked(varargs);SYSCALLS.varargs=varargs;try{var stream=SYSCALLS.getStreamFromFD(fd);switch(cmd){case 0:{var arg=syscallGetVarargI();if(arg<0){return-28}while(FS.streams[arg]){arg++}var newStream;newStream=FS.dupStream(stream,arg);return newStream.fd}case 1:case 2:return 0;case 3:return stream.flags;case 4:{var arg=syscallGetVarargI();stream.flags|=arg;return 0}case 5:{var arg=syscallGetVarargP();var offset=0;(growMemViews(),HEAP16)[(arg+offset)/2]=2;return 0}case 6:case 7:return 0}return-28}catch(e){if(typeof FS=="undefined"||!(e.name==="ErrnoError"))throw e;return-e.errno}}function ___syscall_fstat64(fd,buf){if(ENVIRONMENT_IS_PTHREAD)return proxyToMainThread(4,0,1,fd,buf);buf=bigintToI53Checked(buf);try{return SYSCALLS.writeStat(buf,FS.fstat(fd))}catch(e){if(typeof FS=="undefined"||!(e.name==="ErrnoError"))throw e;return-e.errno}}var stringToUTF8=(str,outPtr,maxBytesToWrite)=>stringToUTF8Array(str,(growMemViews(),HEAPU8),outPtr,maxBytesToWrite);function ___syscall_getcwd(buf,size){if(ENVIRONMENT_IS_PTHREAD)return proxyToMainThread(5,0,1,buf,size);buf=bigintToI53Checked(buf);size=bigintToI53Checked(size);try{if(size===0)return-28;var cwd=FS.cwd();var cwdLengthInBytes=lengthBytesUTF8(cwd)+1;if(size<cwdLengthInBytes)return-68;stringToUTF8(cwd,buf,size);return cwdLengthInBytes}catch(e){if(typeof FS=="undefined"||!(e.name==="ErrnoError"))throw e;return-e.errno}}function ___syscall_getdents64(fd,dirp,count){if(ENVIRONMENT_IS_PTHREAD)return proxyToMainThread(6,0,1,fd,dirp,count);dirp=bigintToI53Checked(dirp);count=bigintToI53Checked(count);try{var stream=SYSCALLS.getStreamFromFD(fd);stream.getdents||=FS.readdir(stream.path);var struct_size=280;var pos=0;var off=FS.llseek(stream,0,1);var startIdx=Math.floor(off/struct_size);var endIdx=Math.min(stream.getdents.length,startIdx+Math.floor(count/struct_size));for(var idx=startIdx;idx<endIdx;idx++){var id;var type;var name=stream.getdents[idx];if(name==="."){id=stream.node.id;type=4}else if(name===".."){var lookup=FS.lookupPath(stream.path,{parent:true});id=lookup.node.id;type=4}else{var child;try{child=FS.lookupNode(stream.node,name)}catch(e){if(e?.errno===28){continue}throw e}id=child.id;type=FS.isChrdev(child.mode)?2:FS.isDir(child.mode)?4:FS.isLink(child.mode)?10:8}(growMemViews(),HEAP64)[(dirp+pos)/8]=BigInt(id);(growMemViews(),HEAP64)[(dirp+pos+8)/8]=BigInt((idx+1)*struct_size);(growMemViews(),HEAP16)[(dirp+pos+16)/2]=280;(growMemViews(),HEAP8)[dirp+pos+18]=type;stringToUTF8(name,dirp+pos+19,256);pos+=struct_size}FS.llseek(stream,idx*struct_size,0);return pos}catch(e){if(typeof FS=="undefined"||!(e.name==="ErrnoError"))throw e;return-e.errno}}function ___syscall_ioctl(fd,op,varargs){if(ENVIRONMENT_IS_PTHREAD)return proxyToMainThread(7,0,1,fd,op,varargs);varargs=bigintToI53Checked(varargs);SYSCALLS.varargs=varargs;try{var stream=SYSCALLS.getStreamFromFD(fd);switch(op){case 21509:{if(!stream.tty)return-59;return 0}case 21505:{if(!stream.tty)return-59;if(stream.tty.ops.ioctl_tcgets){var termios=stream.tty.ops.ioctl_tcgets(stream);var argp=syscallGetVarargP();(growMemViews(),HEAP32)[argp/4]=termios.c_iflag||0;(growMemViews(),HEAP32)[(argp+4)/4]=termios.c_oflag||0;(growMemViews(),HEAP32)[(argp+8)/4]=termios.c_cflag||0;(growMemViews(),HEAP32)[(argp+12)/4]=termios.c_lflag||0;for(var i=0;i<32;i++){(growMemViews(),HEAP8)[argp+i+17]=termios.c_cc[i]||0}return 0}return 0}case 21510:case 21511:case 21512:{if(!stream.tty)return-59;return 0}case 21506:case 21507:case 21508:{if(!stream.tty)return-59;if(stream.tty.ops.ioctl_tcsets){var argp=syscallGetVarargP();var c_iflag=(growMemViews(),HEAP32)[argp/4];var c_oflag=(growMemViews(),HEAP32)[(argp+4)/4];var c_cflag=(growMemViews(),HEAP32)[(argp+8)/4];var c_lflag=(growMemViews(),HEAP32)[(argp+12)/4];var c_cc=[];for(var i=0;i<32;i++){c_cc.push((growMemViews(),HEAP8)[argp+i+17])}return stream.tty.ops.ioctl_tcsets(stream.tty,op,{c_iflag,c_oflag,c_cflag,c_lflag,c_cc})}return 0}case 21519:{if(!stream.tty)return-59;var argp=syscallGetVarargP();(growMemViews(),HEAP32)[argp/4]=0;return 0}case 21520:{if(!stream.tty)return-59;return-28}case 21537:case 21531:{var argp=syscallGetVarargP();return FS.ioctl(stream,op,argp)}case 21523:{if(!stream.tty)return-59;if(stream.tty.ops.ioctl_tiocgwinsz){var winsize=stream.tty.ops.ioctl_tiocgwinsz(stream.tty);var argp=syscallGetVarargP();(growMemViews(),HEAP16)[argp/2]=winsize[0];(growMemViews(),HEAP16)[(argp+2)/2]=winsize[1]}return 0}case 21524:{if(!stream.tty)return-59;return 0}case 21515:{if(!stream.tty)return-59;return 0}default:return-28}}catch(e){if(typeof FS=="undefined"||!(e.name==="ErrnoError"))throw e;return-e.errno}}function ___syscall_lstat64(path,buf){if(ENVIRONMENT_IS_PTHREAD)return proxyToMainThread(8,0,1,path,buf);path=bigintToI53Checked(path);buf=bigintToI53Checked(buf);try{path=SYSCALLS.getStr(path);return SYSCALLS.writeStat(buf,FS.lstat(path))}catch(e){if(typeof FS=="undefined"||!(e.name==="ErrnoError"))throw e;return-e.errno}}function ___syscall_newfstatat(dirfd,path,buf,flags){if(ENVIRONMENT_IS_PTHREAD)return proxyToMainThread(9,0,1,dirfd,path,buf,flags);path=bigintToI53Checked(path);buf=bigintToI53Checked(buf);try{path=SYSCALLS.getStr(path);var nofollow=flags&256;var allowEmpty=flags&4096;flags=flags&~6400;path=SYSCALLS.calculateAt(dirfd,path,allowEmpty);return SYSCALLS.writeStat(buf,nofollow?FS.lstat(path):FS.stat(path))}catch(e){if(typeof FS=="undefined"||!(e.name==="ErrnoError"))throw e;return-e.errno}}function ___syscall_openat(dirfd,path,flags,varargs){if(ENVIRONMENT_IS_PTHREAD)return proxyToMainThread(10,0,1,dirfd,path,flags,varargs);path=bigintToI53Checked(path);varargs=bigintToI53Checked(varargs);SYSCALLS.varargs=varargs;try{path=SYSCALLS.getStr(path);path=SYSCALLS.calculateAt(dirfd,path);var mode=varargs?syscallGetVarargI():0;return FS.open(path,flags,mode).fd}catch(e){if(typeof FS=="undefined"||!(e.name==="ErrnoError"))throw e;return-e.errno}}function ___syscall_stat64(path,buf){if(ENVIRONMENT_IS_PTHREAD)return proxyToMainThread(11,0,1,path,buf);path=bigintToI53Checked(path);buf=bigintToI53Checked(buf);try{path=SYSCALLS.getStr(path);return SYSCALLS.writeStat(buf,FS.stat(path))}catch(e){if(typeof FS=="undefined"||!(e.name==="ErrnoError"))throw e;return-e.errno}}var __abort_js=()=>abort("");function __emscripten_init_main_thread_js(tb){tb=bigintToI53Checked(tb);__emscripten_thread_init(tb,!ENVIRONMENT_IS_WORKER,1,!ENVIRONMENT_IS_WEB,5242880,false);PThread.threadInitTLS()}var handleException=e=>{if(e instanceof ExitStatus||e=="unwind"){return EXITSTATUS}quit_(1,e)};var maybeExit=()=>{if(!keepRuntimeAlive()){try{if(ENVIRONMENT_IS_PTHREAD){if(_pthread_self())__emscripten_thread_exit(EXITSTATUS);return}_exit(EXITSTATUS)}catch(e){handleException(e)}}};var callUserCallback=func=>{if(ABORT){return}try{func();maybeExit()}catch(e){handleException(e)}};function __emscripten_thread_mailbox_await(pthread_ptr){pthread_ptr=bigintToI53Checked(pthread_ptr);if(Atomics.waitAsync){var wait=Atomics.waitAsync((growMemViews(),HEAP32),pthread_ptr/4,pthread_ptr);wait.value.then(checkMailbox);var waitingAsync=pthread_ptr+228;Atomics.store((growMemViews(),HEAP32),waitingAsync/4,1)}}var checkMailbox=()=>callUserCallback(()=>{var pthread_ptr=_pthread_self();if(pthread_ptr){__emscripten_thread_mailbox_await(pthread_ptr);__emscripten_check_mailbox()}});function __emscripten_notify_mailbox_postmessage(targetThread,currThreadId){targetThread=bigintToI53Checked(targetThread);currThreadId=bigintToI53Checked(currThreadId);if(targetThread==currThreadId){setTimeout(checkMailbox)}else if(ENVIRONMENT_IS_PTHREAD){postMessage({targetThread,cmd:"checkMailbox"})}else{var worker=PThread.pthreads[targetThread];if(!worker){return}worker.postMessage({cmd:"checkMailbox"})}}var proxiedJSCallArgs=[];function __emscripten_receive_on_main_thread_js(funcIndex,emAsmAddr,callingThread,numCallArgs,args){emAsmAddr=bigintToI53Checked(emAsmAddr);callingThread=bigintToI53Checked(callingThread);args=bigintToI53Checked(args);numCallArgs/=2;proxiedJSCallArgs.length=numCallArgs;var b=args/8;for(var i=0;i<numCallArgs;i++){if((growMemViews(),HEAP64)[b+2*i]){proxiedJSCallArgs[i]=(growMemViews(),HEAP64)[b+2*i+1]}else{proxiedJSCallArgs[i]=(growMemViews(),HEAPF64)[b+2*i+1]}}var func=proxiedFunctionTable[funcIndex];PThread.currentProxiedOperationCallerThread=callingThread;var rtn=func(...proxiedJSCallArgs);PThread.currentProxiedOperationCallerThread=0;if(typeof rtn=="bigint"){rtn=bigintToI53Checked(rtn)}return rtn}function __emscripten_thread_cleanup(thread){thread=bigintToI53Checked(thread);if(!ENVIRONMENT_IS_PTHREAD)cleanupThread(thread);else postMessage({cmd:"cleanupThread",thread})}function __emscripten_thread_set_strongref(thread){thread=bigintToI53Checked(thread);if(ENVIRONMENT_IS_NODE){PThread.pthreads[thread].ref()}}var isLeapYear=year=>year%4===0&&(year%100!==0||year%400===0);var MONTH_DAYS_LEAP_CUMULATIVE=[0,31,60,91,121,152,182,213,244,274,305,335];var MONTH_DAYS_REGULAR_CUMULATIVE=[0,31,59,90,120,151,181,212,243,273,304,334];var ydayFromDate=date=>{var leap=isLeapYear(date.getFullYear());var monthDaysCumulative=leap?MONTH_DAYS_LEAP_CUMULATIVE:MONTH_DAYS_REGULAR_CUMULATIVE;var yday=monthDaysCumulative[date.getMonth()]+date.getDate()-1;return yday};function __localtime_js(time,tmPtr){time=bigintToI53Checked(time);tmPtr=bigintToI53Checked(tmPtr);var date=new Date(time*1e3);(growMemViews(),HEAP32)[tmPtr/4]=date.getSeconds();(growMemViews(),HEAP32)[(tmPtr+4)/4]=date.getMinutes();(growMemViews(),HEAP32)[(tmPtr+8)/4]=date.getHours();(growMemViews(),HEAP32)[(tmPtr+12)/4]=date.getDate();(growMemViews(),HEAP32)[(tmPtr+16)/4]=date.getMonth();(growMemViews(),HEAP32)[(tmPtr+20)/4]=date.getFullYear()-1900;(growMemViews(),HEAP32)[(tmPtr+24)/4]=date.getDay();var yday=ydayFromDate(date)|0;(growMemViews(),HEAP32)[(tmPtr+28)/4]=yday;(growMemViews(),HEAP64)[(tmPtr+40)/8]=BigInt(-(date.getTimezoneOffset()*60));var start=new Date(date.getFullYear(),0,1);var summerOffset=new Date(date.getFullYear(),6,1).getTimezoneOffset();var winterOffset=start.getTimezoneOffset();var dst=(summerOffset!=winterOffset&&date.getTimezoneOffset()==Math.min(winterOffset,summerOffset))|0;(growMemViews(),HEAP32)[(tmPtr+32)/4]=dst}function __mmap_js(len,prot,flags,fd,offset,allocated,addr){if(ENVIRONMENT_IS_PTHREAD)return proxyToMainThread(12,0,1,len,prot,flags,fd,offset,allocated,addr);len=bigintToI53Checked(len);offset=bigintToI53Checked(offset);allocated=bigintToI53Checked(allocated);addr=bigintToI53Checked(addr);try{var stream=SYSCALLS.getStreamFromFD(fd);var res=FS.mmap(stream,len,offset,prot,flags);var ptr=res.ptr;(growMemViews(),HEAP32)[allocated/4]=res.allocated;(growMemViews(),HEAPU64)[addr/8]=BigInt(ptr);return 0}catch(e){if(typeof FS=="undefined"||!(e.name==="ErrnoError"))throw e;return-e.errno}}function __munmap_js(addr,len,prot,flags,fd,offset){if(ENVIRONMENT_IS_PTHREAD)return proxyToMainThread(13,0,1,addr,len,prot,flags,fd,offset);addr=bigintToI53Checked(addr);len=bigintToI53Checked(len);offset=bigintToI53Checked(offset);try{var stream=SYSCALLS.getStreamFromFD(fd);if(prot&2){SYSCALLS.doMsync(addr,stream,len,flags,offset)}}catch(e){if(typeof FS=="undefined"||!(e.name==="ErrnoError"))throw e;return-e.errno}}var __tzset_js=function(timezone,daylight,std_name,dst_name){timezone=bigintToI53Checked(timezone);daylight=bigintToI53Checked(daylight);std_name=bigintToI53Checked(std_name);dst_name=bigintToI53Checked(dst_name);var currentYear=(new Date).getFullYear();var winter=new Date(currentYear,0,1);var summer=new Date(currentYear,6,1);var winterOffset=winter.getTimezoneOffset();var summerOffset=summer.getTimezoneOffset();var stdTimezoneOffset=Math.max(winterOffset,summerOffset);(growMemViews(),HEAPU64)[timezone/8]=BigInt(stdTimezoneOffset*60);(growMemViews(),HEAP32)[daylight/4]=Number(winterOffset!=summerOffset);var extractZone=timezoneOffset=>{var sign=timezoneOffset>=0?"-":"+";var absOffset=Math.abs(timezoneOffset);var hours=String(Math.floor(absOffset/60)).padStart(2,"0");var minutes=String(absOffset%60).padStart(2,"0");return`UTC${sign}${hours}${minutes}`};var winterName=extractZone(winterOffset);var summerName=extractZone(summerOffset);if(summerOffset<winterOffset){stringToUTF8(winterName,std_name,17);stringToUTF8(summerName,dst_name,17)}else{stringToUTF8(winterName,dst_name,17);stringToUTF8(summerName,std_name,17)}};var _emscripten_get_now=()=>performance.timeOrigin+performance.now();var _emscripten_date_now=()=>Date.now();var nowIsMonotonic=1;var checkWasiClock=clock_id=>clock_id>=0&&clock_id<=3;function _clock_time_get(clk_id,ignored_precision,ptime){ignored_precision=bigintToI53Checked(ignored_precision);ptime=bigintToI53Checked(ptime);if(!checkWasiClock(clk_id)){return 28}var now;if(clk_id===0){now=_emscripten_date_now()}else if(nowIsMonotonic){now=_emscripten_get_now()}else{return 52}var nsec=Math.round(now*1e3*1e3);(growMemViews(),HEAP64)[ptime/8]=BigInt(nsec);return 0}var _emscripten_check_blocking_allowed=()=>{};var runtimeKeepalivePush=()=>{runtimeKeepaliveCounter+=1};var _emscripten_exit_with_live_runtime=()=>{runtimeKeepalivePush();throw"unwind"};var jsStackTrace=()=>(new Error).stack.toString();var getCallstack=flags=>{var callstack=jsStackTrace();var lines=callstack.split("\\n");callstack="";var firefoxRe=new RegExp("\\\\s*(.*?)@(.*?):([0-9]+):([0-9]+)");var chromeRe=new RegExp("\\\\s*at (.*?) \\\\((.*):(.*):(.*)\\\\)");for(var line of lines){var symbolName="";var file="";var lineno=0;var column=0;var parts=chromeRe.exec(line);if(parts?.length==5){symbolName=parts[1];file=parts[2];lineno=parts[3];column=parts[4]}else{parts=firefoxRe.exec(line);if(parts?.length>=4){symbolName=parts[1];file=parts[2];lineno=parts[3];column=parts[4]|0}else{callstack+=line+"\\n";continue}}if(symbolName=="_emscripten_log"||symbolName=="_emscripten_get_callstack"){callstack="";continue}if(flags&24){if(flags&64){file=file.substring(file.replace(/\\\\/g,"/").lastIndexOf("/")+1)}callstack+=`    at ${symbolName} (${file}:${lineno}:${column})\\n`}}callstack=callstack.replace(/\\s+$/,"");return callstack};function _emscripten_get_callstack(flags,str,maxbytes){str=bigintToI53Checked(str);var callstack=getCallstack(flags);if(!str||maxbytes<=0){return lengthBytesUTF8(callstack)+1}var bytesWrittenExcludingNull=stringToUTF8(callstack,str,maxbytes);return bytesWrittenExcludingNull+1}var getHeapMax=()=>4294967296;var _emscripten_get_heap_max=()=>BigInt(getHeapMax());var _emscripten_has_asyncify=()=>2;var _emscripten_num_logical_cores=()=>ENVIRONMENT_IS_NODE?require("os").cpus().length:navigator["hardwareConcurrency"];var growMemory=size=>{var oldHeapSize=wasmMemory.buffer.byteLength;var pages=(size-oldHeapSize+65535)/65536|0;try{wasmMemory.grow(BigInt(pages));updateMemoryViews();return 1}catch(e){}};function _emscripten_resize_heap(requestedSize){requestedSize=bigintToI53Checked(requestedSize);var oldSize=(growMemViews(),HEAPU8).length;if(requestedSize<=oldSize){return false}var maxHeapSize=getHeapMax();if(requestedSize>maxHeapSize){return false}for(var cutDown=1;cutDown<=4;cutDown*=2){var overGrownHeapSize=oldSize*(1+.2/cutDown);overGrownHeapSize=Math.min(overGrownHeapSize,requestedSize+100663296);var newSize=Math.min(maxHeapSize,alignMemory(Math.max(requestedSize,overGrownHeapSize),65536));var replacement=growMemory(newSize);if(replacement){return true}}return false}var stringToUTF8OnStack=str=>{var size=lengthBytesUTF8(str)+1;var ret=stackAlloc(size);stringToUTF8(str,ret,size);return ret};var writeI53ToI64=(ptr,num)=>{(growMemViews(),HEAPU32)[ptr/4]=num;var lower=(growMemViews(),HEAPU32)[ptr/4];(growMemViews(),HEAPU32)[(ptr+4)/4]=(num-lower)/4294967296};var stringToNewUTF8=str=>{var size=lengthBytesUTF8(str)+1;var ret=_malloc(size);if(ret)stringToUTF8(str,ret,size);return ret};var readI53FromI64=ptr=>(growMemViews(),HEAPU32)[ptr/4]+(growMemViews(),HEAP32)[(ptr+4)/4]*4294967296;var WebGPU={Internals:{jsObjects:[],jsObjectInsert:(ptr,jsObject)=>{WebGPU.Internals.jsObjects[ptr]=jsObject},bufferOnUnmaps:[],futures:[],futureInsert:(futureId,promise)=>{WebGPU.Internals.futures[futureId]=new Promise(resolve=>promise.finally(()=>resolve(futureId)))}},getJsObject:ptr=>{if(!ptr)return undefined;return WebGPU.Internals.jsObjects[ptr]},importJsAdapter:(obj,parentPtr=0)=>{var ptr=_emwgpuCreateAdapter(parentPtr);WebGPU.Internals.jsObjects[ptr]=obj;return ptr},importJsBindGroup:(obj,parentPtr=0)=>{var ptr=_emwgpuCreateBindGroup(parentPtr);WebGPU.Internals.jsObjects[ptr]=obj;return ptr},importJsBindGroupLayout:(obj,parentPtr=0)=>{var ptr=_emwgpuCreateBindGroupLayout(parentPtr);WebGPU.Internals.jsObjects[ptr]=obj;return ptr},importJsBuffer:(buffer,parentPtr=0)=>{assert(buffer.mapState==="unmapped");var bufferPtr=_emwgpuCreateBuffer(parentPtr);WebGPU.Internals.jsObjectInsert(bufferPtr,buffer);return bufferPtr},importJsCommandBuffer:(obj,parentPtr=0)=>{var ptr=_emwgpuCreateCommandBuffer(parentPtr);WebGPU.Internals.jsObjects[ptr]=obj;return ptr},importJsCommandEncoder:(obj,parentPtr=0)=>{var ptr=_emwgpuCreateCommandEncoder(parentPtr);WebGPU.Internals.jsObjects[ptr]=obj;return ptr},importJsComputePassEncoder:(obj,parentPtr=0)=>{var ptr=_emwgpuCreateComputePassEncoder(parentPtr);WebGPU.Internals.jsObjects[ptr]=obj;return ptr},importJsComputePipeline:(obj,parentPtr=0)=>{var ptr=_emwgpuCreateComputePipeline(parentPtr);WebGPU.Internals.jsObjects[ptr]=obj;return ptr},importJsDevice:(device,parentPtr=0)=>{var queuePtr=_emwgpuCreateQueue(parentPtr);var devicePtr=_emwgpuCreateDevice(parentPtr,queuePtr);WebGPU.Internals.jsObjectInsert(queuePtr,device.queue);WebGPU.Internals.jsObjectInsert(devicePtr,device);return devicePtr},importJsExternalTexture:(obj,parentPtr=0)=>{var ptr=_emwgpuCreateExternalTexture(parentPtr);WebGPU.Internals.jsObjects[ptr]=obj;return ptr},importJsPipelineLayout:(obj,parentPtr=0)=>{var ptr=_emwgpuCreatePipelineLayout(parentPtr);WebGPU.Internals.jsObjects[ptr]=obj;return ptr},importJsQuerySet:(obj,parentPtr=0)=>{var ptr=_emwgpuCreateQuerySet(parentPtr);WebGPU.Internals.jsObjects[ptr]=obj;return ptr},importJsQueue:(obj,parentPtr=0)=>{var ptr=_emwgpuCreateQueue(parentPtr);WebGPU.Internals.jsObjects[ptr]=obj;return ptr},importJsRenderBundle:(obj,parentPtr=0)=>{var ptr=_emwgpuCreateRenderBundle(parentPtr);WebGPU.Internals.jsObjects[ptr]=obj;return ptr},importJsRenderBundleEncoder:(obj,parentPtr=0)=>{var ptr=_emwgpuCreateRenderBundleEncoder(parentPtr);WebGPU.Internals.jsObjects[ptr]=obj;return ptr},importJsRenderPassEncoder:(obj,parentPtr=0)=>{var ptr=_emwgpuCreateRenderPassEncoder(parentPtr);WebGPU.Internals.jsObjects[ptr]=obj;return ptr},importJsRenderPipeline:(obj,parentPtr=0)=>{var ptr=_emwgpuCreateRenderPipeline(parentPtr);WebGPU.Internals.jsObjects[ptr]=obj;return ptr},importJsSampler:(obj,parentPtr=0)=>{var ptr=_emwgpuCreateSampler(parentPtr);WebGPU.Internals.jsObjects[ptr]=obj;return ptr},importJsShaderModule:(obj,parentPtr=0)=>{var ptr=_emwgpuCreateShaderModule(parentPtr);WebGPU.Internals.jsObjects[ptr]=obj;return ptr},importJsSurface:(obj,parentPtr=0)=>{var ptr=_emwgpuCreateSurface(parentPtr);WebGPU.Internals.jsObjects[ptr]=obj;return ptr},importJsTexture:(obj,parentPtr=0)=>{var ptr=_emwgpuCreateTexture(parentPtr);WebGPU.Internals.jsObjects[ptr]=obj;return ptr},importJsTextureView:(obj,parentPtr=0)=>{var ptr=_emwgpuCreateTextureView(parentPtr);WebGPU.Internals.jsObjects[ptr]=obj;return ptr},errorCallback:(callback,type,message,userdata)=>{var sp=stackSave();var messagePtr=stringToUTF8OnStack(message);((a1,a2,a3)=>getWasmTableEntry(callback).call(null,a1,BigInt(a2),BigInt(a3)))(type,BigInt(messagePtr),userdata);stackRestore(sp)},iterateExtensions:(root,handlers)=>{for(var ptr=Number((growMemViews(),HEAPU64)[root/8]);ptr;ptr=Number((growMemViews(),HEAPU64)[ptr/8])){var sType=(growMemViews(),HEAP32)[(ptr+8)/4];var handler=handlers[sType](ptr)}},setStringView:(ptr,data,length)=>{(growMemViews(),HEAPU64)[ptr/8]=BigInt(data);(growMemViews(),HEAPU64)[(ptr+8)/8]=BigInt(length)},makeStringFromStringView:stringViewPtr=>{var ptr=Number((growMemViews(),HEAPU64)[stringViewPtr/8]);var length=Number((growMemViews(),HEAPU64)[(stringViewPtr+8)/8]);return UTF8ToString(ptr,length)},makeStringFromOptionalStringView:stringViewPtr=>{var ptr=Number((growMemViews(),HEAPU64)[stringViewPtr/8]);var length=Number((growMemViews(),HEAPU64)[(stringViewPtr+8)/8]);if(!ptr){if(length===0){return""}return undefined}return UTF8ToString(ptr,length)},makeColor:ptr=>({r:(growMemViews(),HEAPF64)[ptr/8],g:(growMemViews(),HEAPF64)[(ptr+8)/8],b:(growMemViews(),HEAPF64)[(ptr+16)/8],a:(growMemViews(),HEAPF64)[(ptr+24)/8]}),makeExtent3D:ptr=>({width:(growMemViews(),HEAPU32)[ptr/4],height:(growMemViews(),HEAPU32)[(ptr+4)/4],depthOrArrayLayers:(growMemViews(),HEAPU32)[(ptr+8)/4]}),makeOrigin3D:ptr=>({x:(growMemViews(),HEAPU32)[ptr/4],y:(growMemViews(),HEAPU32)[(ptr+4)/4],z:(growMemViews(),HEAPU32)[(ptr+8)/4]}),makeTexelCopyTextureInfo:ptr=>({texture:WebGPU.getJsObject(Number((growMemViews(),HEAPU64)[ptr/8])),mipLevel:(growMemViews(),HEAPU32)[(ptr+8)/4],origin:WebGPU.makeOrigin3D(ptr+12),aspect:WebGPU.TextureAspect[(growMemViews(),HEAP32)[(ptr+24)/4]]}),makeTexelCopyBufferLayout:ptr=>{var bytesPerRow=(growMemViews(),HEAPU32)[(ptr+8)/4];var rowsPerImage=(growMemViews(),HEAPU32)[(ptr+12)/4];return{offset:readI53FromI64(ptr),bytesPerRow:bytesPerRow===4294967295?undefined:bytesPerRow,rowsPerImage:rowsPerImage===4294967295?undefined:rowsPerImage}},makeTexelCopyBufferInfo:ptr=>{var layoutPtr=ptr+0;var bufferCopyView=WebGPU.makeTexelCopyBufferLayout(layoutPtr);bufferCopyView["buffer"]=WebGPU.getJsObject(Number((growMemViews(),HEAPU64)[(ptr+16)/8]));return bufferCopyView},makePassTimestampWrites:ptr=>{if(ptr===0)return undefined;return{querySet:WebGPU.getJsObject(Number((growMemViews(),HEAPU64)[(ptr+8)/8])),beginningOfPassWriteIndex:(growMemViews(),HEAPU32)[(ptr+16)/4],endOfPassWriteIndex:(growMemViews(),HEAPU32)[(ptr+20)/4]}},makePipelineConstants:(constantCount,constantsPtr)=>{if(!constantCount)return;var constants={};for(var i=0;i<constantCount;++i){var entryPtr=constantsPtr+32*i;var key=WebGPU.makeStringFromStringView(entryPtr+8);constants[key]=(growMemViews(),HEAPF64)[(entryPtr+24)/8]}return constants},makePipelineLayout:layoutPtr=>{if(!layoutPtr)return"auto";return WebGPU.getJsObject(layoutPtr)},makeComputeState:ptr=>{if(!ptr)return undefined;var desc={module:WebGPU.getJsObject(Number((growMemViews(),HEAPU64)[(ptr+8)/8])),constants:WebGPU.makePipelineConstants(Number((growMemViews(),HEAPU64)[(ptr+32)/8]),Number((growMemViews(),HEAPU64)[(ptr+40)/8])),entryPoint:WebGPU.makeStringFromOptionalStringView(ptr+16)};return desc},makeComputePipelineDesc:descriptor=>{var desc={label:WebGPU.makeStringFromOptionalStringView(descriptor+8),layout:WebGPU.makePipelineLayout(Number((growMemViews(),HEAPU64)[(descriptor+24)/8])),compute:WebGPU.makeComputeState(descriptor+32)};return desc},makeRenderPipelineDesc:descriptor=>{function makePrimitiveState(psPtr){if(!psPtr)return undefined;return{topology:WebGPU.PrimitiveTopology[(growMemViews(),HEAP32)[(psPtr+8)/4]],stripIndexFormat:WebGPU.IndexFormat[(growMemViews(),HEAP32)[(psPtr+12)/4]],frontFace:WebGPU.FrontFace[(growMemViews(),HEAP32)[(psPtr+16)/4]],cullMode:WebGPU.CullMode[(growMemViews(),HEAP32)[(psPtr+20)/4]],unclippedDepth:!!(growMemViews(),HEAPU32)[(psPtr+24)/4]}}function makeBlendComponent(bdPtr){if(!bdPtr)return undefined;return{operation:WebGPU.BlendOperation[(growMemViews(),HEAP32)[bdPtr/4]],srcFactor:WebGPU.BlendFactor[(growMemViews(),HEAP32)[(bdPtr+4)/4]],dstFactor:WebGPU.BlendFactor[(growMemViews(),HEAP32)[(bdPtr+8)/4]]}}function makeBlendState(bsPtr){if(!bsPtr)return undefined;return{alpha:makeBlendComponent(bsPtr+12),color:makeBlendComponent(bsPtr+0)}}function makeColorState(csPtr){var format=WebGPU.TextureFormat[(growMemViews(),HEAP32)[(csPtr+8)/4]];return format?{format,blend:makeBlendState(Number((growMemViews(),HEAPU64)[(csPtr+16)/8])),writeMask:(growMemViews(),HEAPU32)[(csPtr+24)/4]}:undefined}function makeColorStates(count,csArrayPtr){var states=[];for(var i=0;i<count;++i){states.push(makeColorState(csArrayPtr+32*i))}return states}function makeStencilStateFace(ssfPtr){return{compare:WebGPU.CompareFunction[(growMemViews(),HEAP32)[ssfPtr/4]],failOp:WebGPU.StencilOperation[(growMemViews(),HEAP32)[(ssfPtr+4)/4]],depthFailOp:WebGPU.StencilOperation[(growMemViews(),HEAP32)[(ssfPtr+8)/4]],passOp:WebGPU.StencilOperation[(growMemViews(),HEAP32)[(ssfPtr+12)/4]]}}function makeDepthStencilState(dssPtr){if(!dssPtr)return undefined;return{format:WebGPU.TextureFormat[(growMemViews(),HEAP32)[(dssPtr+8)/4]],depthWriteEnabled:!!(growMemViews(),HEAPU32)[(dssPtr+12)/4],depthCompare:WebGPU.CompareFunction[(growMemViews(),HEAP32)[(dssPtr+16)/4]],stencilFront:makeStencilStateFace(dssPtr+20),stencilBack:makeStencilStateFace(dssPtr+36),stencilReadMask:(growMemViews(),HEAPU32)[(dssPtr+52)/4],stencilWriteMask:(growMemViews(),HEAPU32)[(dssPtr+56)/4],depthBias:(growMemViews(),HEAP32)[(dssPtr+60)/4],depthBiasSlopeScale:(growMemViews(),HEAPF32)[(dssPtr+64)/4],depthBiasClamp:(growMemViews(),HEAPF32)[(dssPtr+68)/4]}}function makeVertexAttribute(vaPtr){return{format:WebGPU.VertexFormat[(growMemViews(),HEAP32)[(vaPtr+8)/4]],offset:readI53FromI64(vaPtr+16),shaderLocation:(growMemViews(),HEAPU32)[(vaPtr+24)/4]}}function makeVertexAttributes(count,vaArrayPtr){var vas=[];for(var i=0;i<count;++i){vas.push(makeVertexAttribute(vaArrayPtr+i*32))}return vas}function makeVertexBuffer(vbPtr){if(!vbPtr)return undefined;var stepMode=WebGPU.VertexStepMode[(growMemViews(),HEAP32)[(vbPtr+8)/4]];var attributeCount=Number((growMemViews(),HEAPU64)[(vbPtr+24)/8]);if(!stepMode&&!attributeCount){return null}return{arrayStride:readI53FromI64(vbPtr+16),stepMode,attributes:makeVertexAttributes(attributeCount,Number((growMemViews(),HEAPU64)[(vbPtr+32)/8]))}}function makeVertexBuffers(count,vbArrayPtr){if(!count)return undefined;var vbs=[];for(var i=0;i<count;++i){vbs.push(makeVertexBuffer(vbArrayPtr+i*40))}return vbs}function makeVertexState(viPtr){if(!viPtr)return undefined;var desc={module:WebGPU.getJsObject(Number((growMemViews(),HEAPU64)[(viPtr+8)/8])),constants:WebGPU.makePipelineConstants(Number((growMemViews(),HEAPU64)[(viPtr+32)/8]),Number((growMemViews(),HEAPU64)[(viPtr+40)/8])),buffers:makeVertexBuffers(Number((growMemViews(),HEAPU64)[(viPtr+48)/8]),Number((growMemViews(),HEAPU64)[(viPtr+56)/8])),entryPoint:WebGPU.makeStringFromOptionalStringView(viPtr+16)};return desc}function makeMultisampleState(msPtr){if(!msPtr)return undefined;return{count:(growMemViews(),HEAPU32)[(msPtr+8)/4],mask:(growMemViews(),HEAPU32)[(msPtr+12)/4],alphaToCoverageEnabled:!!(growMemViews(),HEAPU32)[(msPtr+16)/4]}}function makeFragmentState(fsPtr){if(!fsPtr)return undefined;var desc={module:WebGPU.getJsObject(Number((growMemViews(),HEAPU64)[(fsPtr+8)/8])),constants:WebGPU.makePipelineConstants(Number((growMemViews(),HEAPU64)[(fsPtr+32)/8]),Number((growMemViews(),HEAPU64)[(fsPtr+40)/8])),targets:makeColorStates(Number((growMemViews(),HEAPU64)[(fsPtr+48)/8]),Number((growMemViews(),HEAPU64)[(fsPtr+56)/8])),entryPoint:WebGPU.makeStringFromOptionalStringView(fsPtr+16)};return desc}var desc={label:WebGPU.makeStringFromOptionalStringView(descriptor+8),layout:WebGPU.makePipelineLayout(Number((growMemViews(),HEAPU64)[(descriptor+24)/8])),vertex:makeVertexState(descriptor+32),primitive:makePrimitiveState(descriptor+96),depthStencil:makeDepthStencilState(Number((growMemViews(),HEAPU64)[(descriptor+128)/8])),multisample:makeMultisampleState(descriptor+136),fragment:makeFragmentState(Number((growMemViews(),HEAPU64)[(descriptor+160)/8]))};return desc},fillLimitStruct:(limits,limitsOutPtr)=>{var nextInChainPtr=Number((growMemViews(),HEAPU64)[limitsOutPtr/8]);function setLimitValueU32(name,basePtr,limitOffset,fallbackValue=0){var limitValue=limits[name]??fallbackValue;(growMemViews(),HEAPU32)[(basePtr+limitOffset)/4]=limitValue}function setLimitValueU64(name,basePtr,limitOffset,fallbackValue=0){var limitValue=limits[name]??fallbackValue;writeI53ToI64(basePtr+limitOffset,limitValue)}setLimitValueU32("maxTextureDimension1D",limitsOutPtr,8);setLimitValueU32("maxTextureDimension2D",limitsOutPtr,12);setLimitValueU32("maxTextureDimension3D",limitsOutPtr,16);setLimitValueU32("maxTextureArrayLayers",limitsOutPtr,20);setLimitValueU32("maxBindGroups",limitsOutPtr,24);setLimitValueU32("maxBindGroupsPlusVertexBuffers",limitsOutPtr,28);setLimitValueU32("maxBindingsPerBindGroup",limitsOutPtr,32);setLimitValueU32("maxDynamicUniformBuffersPerPipelineLayout",limitsOutPtr,36);setLimitValueU32("maxDynamicStorageBuffersPerPipelineLayout",limitsOutPtr,40);setLimitValueU32("maxSampledTexturesPerShaderStage",limitsOutPtr,44);setLimitValueU32("maxSamplersPerShaderStage",limitsOutPtr,48);setLimitValueU32("maxStorageBuffersPerShaderStage",limitsOutPtr,52);setLimitValueU32("maxStorageTexturesPerShaderStage",limitsOutPtr,56);setLimitValueU32("maxUniformBuffersPerShaderStage",limitsOutPtr,60);setLimitValueU32("minUniformBufferOffsetAlignment",limitsOutPtr,80);setLimitValueU32("minStorageBufferOffsetAlignment",limitsOutPtr,84);setLimitValueU64("maxUniformBufferBindingSize",limitsOutPtr,64);setLimitValueU64("maxStorageBufferBindingSize",limitsOutPtr,72);setLimitValueU32("maxVertexBuffers",limitsOutPtr,88);setLimitValueU64("maxBufferSize",limitsOutPtr,96);setLimitValueU32("maxVertexAttributes",limitsOutPtr,104);setLimitValueU32("maxVertexBufferArrayStride",limitsOutPtr,108);setLimitValueU32("maxInterStageShaderVariables",limitsOutPtr,112);setLimitValueU32("maxColorAttachments",limitsOutPtr,116);setLimitValueU32("maxColorAttachmentBytesPerSample",limitsOutPtr,120);setLimitValueU32("maxComputeWorkgroupStorageSize",limitsOutPtr,124);setLimitValueU32("maxComputeInvocationsPerWorkgroup",limitsOutPtr,128);setLimitValueU32("maxComputeWorkgroupSizeX",limitsOutPtr,132);setLimitValueU32("maxComputeWorkgroupSizeY",limitsOutPtr,136);setLimitValueU32("maxComputeWorkgroupSizeZ",limitsOutPtr,140);setLimitValueU32("maxComputeWorkgroupsPerDimension",limitsOutPtr,144);setLimitValueU32("maxImmediateSize",limitsOutPtr,148);if(nextInChainPtr!==0){var sType=(growMemViews(),HEAP32)[(nextInChainPtr+8)/4];var compatibilityModeLimitsPtr=nextInChainPtr;setLimitValueU32("maxStorageBuffersInVertexStage",compatibilityModeLimitsPtr,16,limits.maxStorageBuffersPerShaderStage);setLimitValueU32("maxStorageBuffersInFragmentStage",compatibilityModeLimitsPtr,24,limits.maxStorageBuffersPerShaderStage);setLimitValueU32("maxStorageTexturesInVertexStage",compatibilityModeLimitsPtr,20,limits.maxStorageTexturesPerShaderStage);setLimitValueU32("maxStorageTexturesInFragmentStage",compatibilityModeLimitsPtr,28,limits.maxStorageTexturesPerShaderStage)}},fillAdapterInfoStruct:(info,infoStruct)=>{(growMemViews(),HEAPU32)[(infoStruct+88)/4]=info.subgroupMinSize;(growMemViews(),HEAPU32)[(infoStruct+92)/4]=info.subgroupMaxSize;var strs=info.vendor+info.architecture+info.device+info.description;var strPtr=stringToNewUTF8(strs);var vendorLen=lengthBytesUTF8(info.vendor);WebGPU.setStringView(infoStruct+8,strPtr,vendorLen);strPtr+=vendorLen;var architectureLen=lengthBytesUTF8(info.architecture);WebGPU.setStringView(infoStruct+24,strPtr,architectureLen);strPtr+=architectureLen;var deviceLen=lengthBytesUTF8(info.device);WebGPU.setStringView(infoStruct+40,strPtr,deviceLen);strPtr+=deviceLen;var descriptionLen=lengthBytesUTF8(info.description);WebGPU.setStringView(infoStruct+56,strPtr,descriptionLen);strPtr+=descriptionLen;(growMemViews(),HEAP32)[(infoStruct+72)/4]=2;var adapterType=info.isFallbackAdapter?3:4;(growMemViews(),HEAP32)[(infoStruct+76)/4]=adapterType;(growMemViews(),HEAPU32)[(infoStruct+80)/4]=0;(growMemViews(),HEAPU32)[(infoStruct+84)/4]=0},AddressMode:[,"clamp-to-edge","repeat","mirror-repeat"],BlendFactor:[,"zero","one","src","one-minus-src","src-alpha","one-minus-src-alpha","dst","one-minus-dst","dst-alpha","one-minus-dst-alpha","src-alpha-saturated","constant","one-minus-constant","src1","one-minus-src1","src1-alpha","one-minus-src1-alpha"],BlendOperation:[,"add","subtract","reverse-subtract","min","max"],BufferBindingType:[,,"uniform","storage","read-only-storage"],BufferMapState:[,"unmapped","pending","mapped"],CompareFunction:[,"never","less","equal","less-equal","greater","not-equal","greater-equal","always"],CompilationInfoRequestStatus:[,"success","callback-cancelled"],ComponentSwizzle:[,"0","1","r","g","b","a"],CompositeAlphaMode:[,"opaque","premultiplied","unpremultiplied","inherit"],CullMode:[,"none","front","back"],ErrorFilter:[,"validation","out-of-memory","internal"],FeatureLevel:[,"compatibility","core"],FeatureName:{1:"core-features-and-limits",2:"depth-clip-control",3:"depth32float-stencil8",4:"texture-compression-bc",5:"texture-compression-bc-sliced-3d",6:"texture-compression-etc2",7:"texture-compression-astc",8:"texture-compression-astc-sliced-3d",9:"timestamp-query",10:"indirect-first-instance",11:"shader-f16",12:"rg11b10ufloat-renderable",13:"bgra8unorm-storage",14:"float32-filterable",15:"float32-blendable",16:"clip-distances",17:"dual-source-blending",18:"subgroups",19:"texture-formats-tier1",20:"texture-formats-tier2",21:"primitive-index",22:"texture-component-swizzle",327692:"chromium-experimental-unorm16-texture-formats",327729:"chromium-experimental-multi-draw-indirect"},FilterMode:[,"nearest","linear"],FrontFace:[,"ccw","cw"],IndexFormat:[,"uint16","uint32"],InstanceFeatureName:[,"timed-wait-any","shader-source-spirv","multiple-devices-per-adapter"],LoadOp:[,"load","clear"],MipmapFilterMode:[,"nearest","linear"],OptionalBool:["false","true"],PowerPreference:[,"low-power","high-performance"],PredefinedColorSpace:[,"srgb","display-p3"],PrimitiveTopology:[,"point-list","line-list","line-strip","triangle-list","triangle-strip"],QueryType:[,"occlusion","timestamp"],SamplerBindingType:[,,"filtering","non-filtering","comparison"],Status:[,"success","error"],StencilOperation:[,"keep","zero","replace","invert","increment-clamp","decrement-clamp","increment-wrap","decrement-wrap"],StorageTextureAccess:[,,"write-only","read-only","read-write"],StoreOp:[,"store","discard"],SurfaceGetCurrentTextureStatus:[,"success-optimal","success-suboptimal","timeout","outdated","lost","error"],TextureAspect:[,"all","stencil-only","depth-only"],TextureDimension:[,"1d","2d","3d"],TextureFormat:[,"r8unorm","r8snorm","r8uint","r8sint","r16unorm","r16snorm","r16uint","r16sint","r16float","rg8unorm","rg8snorm","rg8uint","rg8sint","r32float","r32uint","r32sint","rg16unorm","rg16snorm","rg16uint","rg16sint","rg16float","rgba8unorm","rgba8unorm-srgb","rgba8snorm","rgba8uint","rgba8sint","bgra8unorm","bgra8unorm-srgb","rgb10a2uint","rgb10a2unorm","rg11b10ufloat","rgb9e5ufloat","rg32float","rg32uint","rg32sint","rgba16unorm","rgba16snorm","rgba16uint","rgba16sint","rgba16float","rgba32float","rgba32uint","rgba32sint","stencil8","depth16unorm","depth24plus","depth24plus-stencil8","depth32float","depth32float-stencil8","bc1-rgba-unorm","bc1-rgba-unorm-srgb","bc2-rgba-unorm","bc2-rgba-unorm-srgb","bc3-rgba-unorm","bc3-rgba-unorm-srgb","bc4-r-unorm","bc4-r-snorm","bc5-rg-unorm","bc5-rg-snorm","bc6h-rgb-ufloat","bc6h-rgb-float","bc7-rgba-unorm","bc7-rgba-unorm-srgb","etc2-rgb8unorm","etc2-rgb8unorm-srgb","etc2-rgb8a1unorm","etc2-rgb8a1unorm-srgb","etc2-rgba8unorm","etc2-rgba8unorm-srgb","eac-r11unorm","eac-r11snorm","eac-rg11unorm","eac-rg11snorm","astc-4x4-unorm","astc-4x4-unorm-srgb","astc-5x4-unorm","astc-5x4-unorm-srgb","astc-5x5-unorm","astc-5x5-unorm-srgb","astc-6x5-unorm","astc-6x5-unorm-srgb","astc-6x6-unorm","astc-6x6-unorm-srgb","astc-8x5-unorm","astc-8x5-unorm-srgb","astc-8x6-unorm","astc-8x6-unorm-srgb","astc-8x8-unorm","astc-8x8-unorm-srgb","astc-10x5-unorm","astc-10x5-unorm-srgb","astc-10x6-unorm","astc-10x6-unorm-srgb","astc-10x8-unorm","astc-10x8-unorm-srgb","astc-10x10-unorm","astc-10x10-unorm-srgb","astc-12x10-unorm","astc-12x10-unorm-srgb","astc-12x12-unorm","astc-12x12-unorm-srgb"],TextureSampleType:[,,"float","unfilterable-float","depth","sint","uint"],TextureViewDimension:[,"1d","2d","2d-array","cube","cube-array","3d"],ToneMappingMode:[,"standard","extended"],VertexFormat:[,"uint8","uint8x2","uint8x4","sint8","sint8x2","sint8x4","unorm8","unorm8x2","unorm8x4","snorm8","snorm8x2","snorm8x4","uint16","uint16x2","uint16x4","sint16","sint16x2","sint16x4","unorm16","unorm16x2","unorm16x4","snorm16","snorm16x2","snorm16x4","float16","float16x2","float16x4","float32","float32x2","float32x3","float32x4","uint32","uint32x2","uint32x3","uint32x4","sint32","sint32x2","sint32x3","sint32x4","unorm10-10-10-2","unorm8x4-bgra"],VertexStepMode:[,"vertex","instance"],WGSLLanguageFeatureName:[,"readonly_and_readwrite_storage_textures","packed_4x8_integer_dot_product","unrestricted_pointer_parameters","pointer_composite_access","uniform_buffer_standard_layout","subgroup_id","texture_and_sampler_let","subgroup_uniformity","texture_formats_tier1"]};var emwgpuStringToInt_DeviceLostReason={undefined:1,unknown:1,destroyed:2};var runtimeKeepalivePop=()=>{runtimeKeepaliveCounter-=1};function _emwgpuAdapterRequestDevice(adapterPtr,futureId,deviceLostFutureId,devicePtr,queuePtr,descriptor){adapterPtr=bigintToI53Checked(adapterPtr);futureId=bigintToI53Checked(futureId);deviceLostFutureId=bigintToI53Checked(deviceLostFutureId);devicePtr=bigintToI53Checked(devicePtr);queuePtr=bigintToI53Checked(queuePtr);descriptor=bigintToI53Checked(descriptor);var adapter=WebGPU.getJsObject(adapterPtr);var desc={};if(descriptor){var requiredFeatureCount=Number((growMemViews(),HEAPU64)[(descriptor+24)/8]);if(requiredFeatureCount){var requiredFeaturesPtr=Number((growMemViews(),HEAPU64)[(descriptor+32)/8]);desc["requiredFeatures"]=Array.from((growMemViews(),HEAPU32).subarray(requiredFeaturesPtr/4,(requiredFeaturesPtr+requiredFeatureCount*4)/4),feature=>WebGPU.FeatureName[feature])}var limitsPtr=Number((growMemViews(),HEAPU64)[(descriptor+40)/8]);if(limitsPtr){var nextInChainPtr=Number((growMemViews(),HEAPU64)[limitsPtr/8]);var requiredLimits={};function setLimitU32IfDefined(name,basePtr,limitOffset,ignoreIfZero=false){var ptr=basePtr+limitOffset;var value=(growMemViews(),HEAPU32)[ptr/4];if(value!=4294967295&&(!ignoreIfZero||value!=0)){requiredLimits[name]=value}}function setLimitU64IfDefined(name,basePtr,limitOffset){var ptr=basePtr+limitOffset;var limitPart1=(growMemViews(),HEAPU32)[ptr/4];var limitPart2=(growMemViews(),HEAPU32)[(ptr+4)/4];if(limitPart1!=4294967295||limitPart2!=4294967295){requiredLimits[name]=readI53FromI64(ptr)}}setLimitU32IfDefined("maxTextureDimension1D",limitsPtr,8);setLimitU32IfDefined("maxTextureDimension2D",limitsPtr,12);setLimitU32IfDefined("maxTextureDimension3D",limitsPtr,16);setLimitU32IfDefined("maxTextureArrayLayers",limitsPtr,20);setLimitU32IfDefined("maxBindGroups",limitsPtr,24);setLimitU32IfDefined("maxBindGroupsPlusVertexBuffers",limitsPtr,28);setLimitU32IfDefined("maxBindingsPerBindGroup",limitsPtr,32);setLimitU32IfDefined("maxDynamicUniformBuffersPerPipelineLayout",limitsPtr,36);setLimitU32IfDefined("maxDynamicStorageBuffersPerPipelineLayout",limitsPtr,40);setLimitU32IfDefined("maxSampledTexturesPerShaderStage",limitsPtr,44);setLimitU32IfDefined("maxSamplersPerShaderStage",limitsPtr,48);setLimitU32IfDefined("maxStorageBuffersPerShaderStage",limitsPtr,52);setLimitU32IfDefined("maxStorageTexturesPerShaderStage",limitsPtr,56);setLimitU32IfDefined("maxUniformBuffersPerShaderStage",limitsPtr,60);setLimitU32IfDefined("minUniformBufferOffsetAlignment",limitsPtr,80);setLimitU32IfDefined("minStorageBufferOffsetAlignment",limitsPtr,84);setLimitU64IfDefined("maxUniformBufferBindingSize",limitsPtr,64);setLimitU64IfDefined("maxStorageBufferBindingSize",limitsPtr,72);setLimitU32IfDefined("maxVertexBuffers",limitsPtr,88);setLimitU64IfDefined("maxBufferSize",limitsPtr,96);setLimitU32IfDefined("maxVertexAttributes",limitsPtr,104);setLimitU32IfDefined("maxVertexBufferArrayStride",limitsPtr,108);setLimitU32IfDefined("maxInterStageShaderVariables",limitsPtr,112);setLimitU32IfDefined("maxColorAttachments",limitsPtr,116);setLimitU32IfDefined("maxColorAttachmentBytesPerSample",limitsPtr,120);setLimitU32IfDefined("maxComputeWorkgroupStorageSize",limitsPtr,124);setLimitU32IfDefined("maxComputeInvocationsPerWorkgroup",limitsPtr,128);setLimitU32IfDefined("maxComputeWorkgroupSizeX",limitsPtr,132);setLimitU32IfDefined("maxComputeWorkgroupSizeY",limitsPtr,136);setLimitU32IfDefined("maxComputeWorkgroupSizeZ",limitsPtr,140);setLimitU32IfDefined("maxComputeWorkgroupsPerDimension",limitsPtr,144);setLimitU32IfDefined("maxImmediateSize",limitsPtr,148,true);if(nextInChainPtr!==0){var sType=(growMemViews(),HEAP32)[(nextInChainPtr+8)/4];var compatibilityModeLimitsPtr=nextInChainPtr;if("maxStorageBuffersInVertexStage"in GPUSupportedLimits.prototype){setLimitU32IfDefined("maxStorageBuffersInVertexStage",compatibilityModeLimitsPtr,16);setLimitU32IfDefined("maxStorageTexturesInVertexStage",compatibilityModeLimitsPtr,20);setLimitU32IfDefined("maxStorageBuffersInFragmentStage",compatibilityModeLimitsPtr,24);setLimitU32IfDefined("maxStorageTexturesInFragmentStage",compatibilityModeLimitsPtr,28)}}desc["requiredLimits"]=requiredLimits}var defaultQueuePtr=Number((growMemViews(),HEAPU64)[(descriptor+48)/8]);if(defaultQueuePtr){var defaultQueueDesc={label:WebGPU.makeStringFromOptionalStringView(defaultQueuePtr+8)};desc["defaultQueue"]=defaultQueueDesc}desc["label"]=WebGPU.makeStringFromOptionalStringView(descriptor+8)}runtimeKeepalivePush();WebGPU.Internals.futureInsert(futureId,adapter.requestDevice(desc).then(device=>{runtimeKeepalivePop();callUserCallback(()=>{WebGPU.Internals.jsObjectInsert(queuePtr,device.queue);WebGPU.Internals.jsObjectInsert(devicePtr,device);devicePtr=BigInt(devicePtr);WebGPU.Internals.futureInsert(deviceLostFutureId,device.lost.then(info=>{callUserCallback(()=>{device.onuncapturederror=ev=>{};var sp=stackSave();var messagePtr=stringToUTF8OnStack(info.message);_emwgpuOnDeviceLostCompleted(deviceLostFutureId,emwgpuStringToInt_DeviceLostReason[info.reason],BigInt(messagePtr));stackRestore(sp)})}));device.onuncapturederror=ev=>{var type=5;if(ev.error instanceof GPUValidationError)type=2;else if(ev.error instanceof GPUOutOfMemoryError)type=3;else if(ev.error instanceof GPUInternalError)type=4;var sp=stackSave();var messagePtr=stringToUTF8OnStack(ev.error.message);_emwgpuOnUncapturedError(BigInt(devicePtr),type,BigInt(messagePtr));stackRestore(sp)};_emwgpuOnRequestDeviceCompleted(futureId,1,BigInt(devicePtr),0n)})},ex=>{runtimeKeepalivePop();callUserCallback(()=>{var sp=stackSave();var messagePtr=stringToUTF8OnStack(ex.message);_emwgpuOnRequestDeviceCompleted(futureId,3,BigInt(devicePtr),BigInt(messagePtr));if(deviceLostFutureId){_emwgpuOnDeviceLostCompleted(deviceLostFutureId,4,BigInt(messagePtr))}stackRestore(sp)})}))}function _emwgpuBufferDestroy(bufferPtr){bufferPtr=bigintToI53Checked(bufferPtr);var buffer=WebGPU.getJsObject(bufferPtr);var onUnmap=WebGPU.Internals.bufferOnUnmaps[bufferPtr];if(onUnmap){for(var i=0;i<onUnmap.length;++i){onUnmap[i]()}delete WebGPU.Internals.bufferOnUnmaps[bufferPtr]}buffer.destroy()}var warnOnce=text=>{warnOnce.shown||={};if(!warnOnce.shown[text]){warnOnce.shown[text]=1;if(ENVIRONMENT_IS_NODE)text="warning: "+text;err(text)}};var _emwgpuBufferGetConstMappedRange=function(bufferPtr,offset,size){bufferPtr=bigintToI53Checked(bufferPtr);offset=bigintToI53Checked(offset);size=bigintToI53Checked(size);var ret=(()=>{var buffer=WebGPU.getJsObject(bufferPtr);if(size==-1)size=undefined;var mapped;try{mapped=buffer.getMappedRange(offset,size)}catch(ex){return 0n}var data=_memalign(16,mapped.byteLength);(growMemViews(),HEAPU8).set(new Uint8Array(mapped),data);WebGPU.Internals.bufferOnUnmaps[bufferPtr].push(()=>_free(data));return data})();return BigInt(ret)};var _emwgpuBufferMapAsync=function(bufferPtr,futureId,mode,offset,size){bufferPtr=bigintToI53Checked(bufferPtr);futureId=bigintToI53Checked(futureId);mode=bigintToI53Checked(mode);offset=bigintToI53Checked(offset);size=bigintToI53Checked(size);var buffer=WebGPU.getJsObject(bufferPtr);WebGPU.Internals.bufferOnUnmaps[bufferPtr]=[];if(size==-1)size=undefined;runtimeKeepalivePush();WebGPU.Internals.futureInsert(futureId,buffer.mapAsync(mode,offset,size).then(()=>{runtimeKeepalivePop();callUserCallback(()=>{_emwgpuOnMapAsyncCompleted(futureId,1,0n)})},ex=>{runtimeKeepalivePop();callUserCallback(()=>{var sp=stackSave();var messagePtr=stringToUTF8OnStack(ex.message);var status=ex.name==="AbortError"?4:ex.name==="OperationError"?3:0;_emwgpuOnMapAsyncCompleted(futureId,status,BigInt(messagePtr));delete WebGPU.Internals.bufferOnUnmaps[bufferPtr]})}))};function _emwgpuBufferUnmap(bufferPtr){bufferPtr=bigintToI53Checked(bufferPtr);var buffer=WebGPU.getJsObject(bufferPtr);var onUnmap=WebGPU.Internals.bufferOnUnmaps[bufferPtr];if(!onUnmap){return}for(var i=0;i<onUnmap.length;++i){onUnmap[i]()}delete WebGPU.Internals.bufferOnUnmaps[bufferPtr];buffer.unmap()}function _emwgpuDelete(ptr){ptr=bigintToI53Checked(ptr);delete WebGPU.Internals.jsObjects[ptr]}function _emwgpuDeviceCreateBuffer(devicePtr,descriptor,bufferPtr){devicePtr=bigintToI53Checked(devicePtr);descriptor=bigintToI53Checked(descriptor);bufferPtr=bigintToI53Checked(bufferPtr);var mappedAtCreation=!!(growMemViews(),HEAPU32)[(descriptor+40)/4];var desc={label:WebGPU.makeStringFromOptionalStringView(descriptor+8),usage:(growMemViews(),HEAPU32)[(descriptor+24)/4],size:readI53FromI64(descriptor+32),mappedAtCreation};var device=WebGPU.getJsObject(devicePtr);var buffer;try{buffer=device.createBuffer(desc)}catch(ex){return false}WebGPU.Internals.jsObjectInsert(bufferPtr,buffer);if(mappedAtCreation){WebGPU.Internals.bufferOnUnmaps[bufferPtr]=[]}return true}function _emwgpuDeviceCreateShaderModule(devicePtr,descriptor,shaderModulePtr){devicePtr=bigintToI53Checked(devicePtr);descriptor=bigintToI53Checked(descriptor);shaderModulePtr=bigintToI53Checked(shaderModulePtr);var nextInChainPtr=Number((growMemViews(),HEAPU64)[descriptor/8]);var sType=(growMemViews(),HEAP32)[(nextInChainPtr+8)/4];var desc={label:WebGPU.makeStringFromOptionalStringView(descriptor+8),code:""};switch(sType){case 2:{desc["code"]=WebGPU.makeStringFromStringView(nextInChainPtr+16);break}}var device=WebGPU.getJsObject(devicePtr);WebGPU.Internals.jsObjectInsert(shaderModulePtr,device.createShaderModule(desc))}var _emwgpuDeviceDestroy=devicePtr=>{const device=WebGPU.getJsObject(devicePtr);device.onuncapturederror=null;device.destroy()};function _emwgpuInstanceRequestAdapter(instancePtr,futureId,options,adapterPtr){instancePtr=bigintToI53Checked(instancePtr);futureId=bigintToI53Checked(futureId);options=bigintToI53Checked(options);adapterPtr=bigintToI53Checked(adapterPtr);var opts;if(options){opts={featureLevel:WebGPU.FeatureLevel[(growMemViews(),HEAP32)[(options+8)/4]],powerPreference:WebGPU.PowerPreference[(growMemViews(),HEAP32)[(options+12)/4]],forceFallbackAdapter:!!(growMemViews(),HEAPU32)[(options+16)/4]};var nextInChainPtr=Number((growMemViews(),HEAPU64)[options/8]);if(nextInChainPtr!==0){var sType=(growMemViews(),HEAP32)[(nextInChainPtr+8)/4];var webxrOptions=nextInChainPtr;opts.xrCompatible=!!(growMemViews(),HEAPU32)[(webxrOptions+16)/4]}}if(!("gpu"in navigator)){var sp=stackSave();var messagePtr=stringToUTF8OnStack("WebGPU not available on this browser (navigator.gpu is not available)");_emwgpuOnRequestAdapterCompleted(futureId,3,BigInt(adapterPtr),BigInt(messagePtr));stackRestore(sp);return}runtimeKeepalivePush();WebGPU.Internals.futureInsert(futureId,navigator.gpu.requestAdapter(opts).then(adapter=>{runtimeKeepalivePop();callUserCallback(()=>{if(adapter){WebGPU.Internals.jsObjectInsert(adapterPtr,adapter);_emwgpuOnRequestAdapterCompleted(futureId,1,BigInt(adapterPtr),0n)}else{var sp=stackSave();var messagePtr=stringToUTF8OnStack("WebGPU not available on this browser (requestAdapter returned null)");_emwgpuOnRequestAdapterCompleted(futureId,3,BigInt(adapterPtr),BigInt(messagePtr));stackRestore(sp)}})},ex=>{runtimeKeepalivePop();callUserCallback(()=>{var sp=stackSave();var messagePtr=stringToUTF8OnStack(ex.message);_emwgpuOnRequestAdapterCompleted(futureId,4,BigInt(adapterPtr),BigInt(messagePtr));stackRestore(sp)})}))}var _emwgpuQueueOnSubmittedWorkDone=function(queuePtr,futureId){queuePtr=bigintToI53Checked(queuePtr);futureId=bigintToI53Checked(futureId);var queue=WebGPU.getJsObject(queuePtr);runtimeKeepalivePush();WebGPU.Internals.futureInsert(futureId,queue.onSubmittedWorkDone().then(()=>{runtimeKeepalivePop();callUserCallback(()=>{_emwgpuOnWorkDoneCompleted(futureId,1)})}))};var _emwgpuWaitAny=function(futurePtr,futureCount,timeoutMSPtr){futurePtr=bigintToI53Checked(futurePtr);futureCount=bigintToI53Checked(futureCount);timeoutMSPtr=bigintToI53Checked(timeoutMSPtr);return Asyncify.handleAsync(async()=>{var promises=[];if(timeoutMSPtr){var timeoutMS=(growMemViews(),HEAP32)[timeoutMSPtr/4];promises.length=futureCount+1;promises[futureCount]=new Promise(resolve=>setTimeout(resolve,timeoutMS,0))}else{promises.length=futureCount}for(var i=0;i<futureCount;++i){var futureId=readI53FromI64(futurePtr+i*8);if(!(futureId in WebGPU.Internals.futures)){return futureId}promises[i]=WebGPU.Internals.futures[futureId]}const firstResolvedFuture=await Promise.race(promises);delete WebGPU.Internals.futures[firstResolvedFuture];return firstResolvedFuture})};_emwgpuWaitAny.isAsync=true;var ENV={};var getExecutableName=()=>thisProgram||"./this.program";var getEnvStrings=()=>{if(!getEnvStrings.strings){var lang=(typeof navigator=="object"&&navigator.language||"C").replace("-","_")+".UTF-8";var env={USER:"web_user",LOGNAME:"web_user",PATH:"/",PWD:"/",HOME:"/home/web_user",LANG:lang,_:getExecutableName()};for(var x in ENV){if(ENV[x]===undefined)delete env[x];else env[x]=ENV[x]}var strings=[];for(var x in env){strings.push(`${x}=${env[x]}`)}getEnvStrings.strings=strings}return getEnvStrings.strings};function _environ_get(__environ,environ_buf){if(ENVIRONMENT_IS_PTHREAD)return proxyToMainThread(14,0,1,__environ,environ_buf);__environ=bigintToI53Checked(__environ);environ_buf=bigintToI53Checked(environ_buf);var bufSize=0;var envp=0;for(var string of getEnvStrings()){var ptr=environ_buf+bufSize;(growMemViews(),HEAPU64)[(__environ+envp)/8]=BigInt(ptr);bufSize+=stringToUTF8(string,ptr,Infinity)+1;envp+=8}return 0}function _environ_sizes_get(penviron_count,penviron_buf_size){if(ENVIRONMENT_IS_PTHREAD)return proxyToMainThread(15,0,1,penviron_count,penviron_buf_size);penviron_count=bigintToI53Checked(penviron_count);penviron_buf_size=bigintToI53Checked(penviron_buf_size);var strings=getEnvStrings();(growMemViews(),HEAPU64)[penviron_count/8]=BigInt(strings.length);var bufSize=0;for(var string of strings){bufSize+=lengthBytesUTF8(string)+1}(growMemViews(),HEAPU64)[penviron_buf_size/8]=BigInt(bufSize);return 0}function _fd_close(fd){if(ENVIRONMENT_IS_PTHREAD)return proxyToMainThread(16,0,1,fd);try{var stream=SYSCALLS.getStreamFromFD(fd);FS.close(stream);return 0}catch(e){if(typeof FS=="undefined"||!(e.name==="ErrnoError"))throw e;return e.errno}}var doReadv=(stream,iov,iovcnt,offset)=>{var ret=0;for(var i=0;i<iovcnt;i++){var ptr=Number((growMemViews(),HEAPU64)[iov/8]);var len=Number((growMemViews(),HEAPU64)[(iov+8)/8]);iov+=16;var curr=FS.read(stream,(growMemViews(),HEAP8),ptr,len,offset);if(curr<0)return-1;ret+=curr;if(curr<len)break;if(typeof offset!="undefined"){offset+=curr}}return ret};function _fd_read(fd,iov,iovcnt,pnum){if(ENVIRONMENT_IS_PTHREAD)return proxyToMainThread(17,0,1,fd,iov,iovcnt,pnum);iov=bigintToI53Checked(iov);iovcnt=bigintToI53Checked(iovcnt);pnum=bigintToI53Checked(pnum);try{var stream=SYSCALLS.getStreamFromFD(fd);var num=doReadv(stream,iov,iovcnt);(growMemViews(),HEAPU64)[pnum/8]=BigInt(num);return 0}catch(e){if(typeof FS=="undefined"||!(e.name==="ErrnoError"))throw e;return e.errno}}function _fd_seek(fd,offset,whence,newOffset){if(ENVIRONMENT_IS_PTHREAD)return proxyToMainThread(18,0,1,fd,offset,whence,newOffset);offset=bigintToI53Checked(offset);newOffset=bigintToI53Checked(newOffset);try{if(isNaN(offset))return 61;var stream=SYSCALLS.getStreamFromFD(fd);FS.llseek(stream,offset,whence);(growMemViews(),HEAP64)[newOffset/8]=BigInt(stream.position);if(stream.getdents&&offset===0&&whence===0)stream.getdents=null;return 0}catch(e){if(typeof FS=="undefined"||!(e.name==="ErrnoError"))throw e;return e.errno}}var doWritev=(stream,iov,iovcnt,offset)=>{var ret=0;for(var i=0;i<iovcnt;i++){var ptr=Number((growMemViews(),HEAPU64)[iov/8]);var len=Number((growMemViews(),HEAPU64)[(iov+8)/8]);iov+=16;var curr=FS.write(stream,(growMemViews(),HEAP8),ptr,len,offset);if(curr<0)return-1;ret+=curr;if(curr<len){break}if(typeof offset!="undefined"){offset+=curr}}return ret};function _fd_write(fd,iov,iovcnt,pnum){if(ENVIRONMENT_IS_PTHREAD)return proxyToMainThread(19,0,1,fd,iov,iovcnt,pnum);iov=bigintToI53Checked(iov);iovcnt=bigintToI53Checked(iovcnt);pnum=bigintToI53Checked(pnum);try{var stream=SYSCALLS.getStreamFromFD(fd);var num=doWritev(stream,iov,iovcnt);(growMemViews(),HEAPU64)[pnum/8]=BigInt(num);return 0}catch(e){if(typeof FS=="undefined"||!(e.name==="ErrnoError"))throw e;return e.errno}}function _random_get(buffer,size){buffer=bigintToI53Checked(buffer);size=bigintToI53Checked(size);try{randomFill((growMemViews(),HEAPU8).subarray(buffer,buffer+size));return 0}catch(e){if(typeof FS=="undefined"||!(e.name==="ErrnoError"))throw e;return e.errno}}function _wgpuAdapterGetInfo(adapterPtr,info){adapterPtr=bigintToI53Checked(adapterPtr);info=bigintToI53Checked(info);var adapter=WebGPU.getJsObject(adapterPtr);WebGPU.fillAdapterInfoStruct(adapter.info,info);return 1}function _wgpuAdapterGetLimits(adapterPtr,limitsOutPtr){adapterPtr=bigintToI53Checked(adapterPtr);limitsOutPtr=bigintToI53Checked(limitsOutPtr);var adapter=WebGPU.getJsObject(adapterPtr);WebGPU.fillLimitStruct(adapter.limits,limitsOutPtr);return 1}function _wgpuAdapterHasFeature(adapterPtr,featureEnumValue){adapterPtr=bigintToI53Checked(adapterPtr);var adapter=WebGPU.getJsObject(adapterPtr);return adapter.features.has(WebGPU.FeatureName[featureEnumValue])}var _wgpuBufferGetSize=function(bufferPtr){bufferPtr=bigintToI53Checked(bufferPtr);var ret=(()=>{var buffer=WebGPU.getJsObject(bufferPtr);return buffer.size})();return BigInt(ret)};var _wgpuCommandEncoderBeginComputePass=function(encoderPtr,descriptor){encoderPtr=bigintToI53Checked(encoderPtr);descriptor=bigintToI53Checked(descriptor);var ret=(()=>{var desc;if(descriptor){desc={label:WebGPU.makeStringFromOptionalStringView(descriptor+8),timestampWrites:WebGPU.makePassTimestampWrites(Number((growMemViews(),HEAPU64)[(descriptor+24)/8]))}}var commandEncoder=WebGPU.getJsObject(encoderPtr);var ptr=_emwgpuCreateComputePassEncoder(0n);WebGPU.Internals.jsObjectInsert(ptr,commandEncoder.beginComputePass(desc));return ptr})();return BigInt(ret)};function _wgpuCommandEncoderCopyBufferToBuffer(encoderPtr,srcPtr,srcOffset,dstPtr,dstOffset,size){encoderPtr=bigintToI53Checked(encoderPtr);srcPtr=bigintToI53Checked(srcPtr);srcOffset=bigintToI53Checked(srcOffset);dstPtr=bigintToI53Checked(dstPtr);dstOffset=bigintToI53Checked(dstOffset);size=bigintToI53Checked(size);var commandEncoder=WebGPU.getJsObject(encoderPtr);var src=WebGPU.getJsObject(srcPtr);var dst=WebGPU.getJsObject(dstPtr);commandEncoder.copyBufferToBuffer(src,srcOffset,dst,dstOffset,size)}var _wgpuCommandEncoderFinish=function(encoderPtr,descriptor){encoderPtr=bigintToI53Checked(encoderPtr);descriptor=bigintToI53Checked(descriptor);var ret=(()=>{var commandEncoder=WebGPU.getJsObject(encoderPtr);var ptr=_emwgpuCreateCommandBuffer(0n);WebGPU.Internals.jsObjectInsert(ptr,commandEncoder.finish());return ptr})();return BigInt(ret)};function _wgpuComputePassEncoderDispatchWorkgroups(passPtr,x,y,z){passPtr=bigintToI53Checked(passPtr);var pass=WebGPU.getJsObject(passPtr);pass.dispatchWorkgroups(x,y,z)}function _wgpuComputePassEncoderEnd(passPtr){passPtr=bigintToI53Checked(passPtr);var pass=WebGPU.getJsObject(passPtr);pass.end()}function _wgpuComputePassEncoderSetBindGroup(passPtr,groupIndex,groupPtr,dynamicOffsetCount,dynamicOffsetsPtr){passPtr=bigintToI53Checked(passPtr);groupPtr=bigintToI53Checked(groupPtr);dynamicOffsetCount=bigintToI53Checked(dynamicOffsetCount);dynamicOffsetsPtr=bigintToI53Checked(dynamicOffsetsPtr);var pass=WebGPU.getJsObject(passPtr);var group=WebGPU.getJsObject(groupPtr);if(dynamicOffsetCount==0){pass.setBindGroup(groupIndex,group)}else{pass.setBindGroup(groupIndex,group,(growMemViews(),HEAPU32),dynamicOffsetsPtr/4,dynamicOffsetCount)}}function _wgpuComputePassEncoderSetPipeline(passPtr,pipelinePtr){passPtr=bigintToI53Checked(passPtr);pipelinePtr=bigintToI53Checked(pipelinePtr);var pass=WebGPU.getJsObject(passPtr);var pipeline=WebGPU.getJsObject(pipelinePtr);pass.setPipeline(pipeline)}var _wgpuComputePipelineGetBindGroupLayout=function(pipelinePtr,groupIndex){pipelinePtr=bigintToI53Checked(pipelinePtr);var ret=(()=>{var pipeline=WebGPU.getJsObject(pipelinePtr);var ptr=_emwgpuCreateBindGroupLayout(0n);WebGPU.Internals.jsObjectInsert(ptr,pipeline.getBindGroupLayout(groupIndex));return ptr})();return BigInt(ret)};var _wgpuDeviceCreateBindGroup=function(devicePtr,descriptor){devicePtr=bigintToI53Checked(devicePtr);descriptor=bigintToI53Checked(descriptor);var ret=(()=>{function makeEntry(entryPtr){var bufferPtr=Number((growMemViews(),HEAPU64)[(entryPtr+16)/8]);var samplerPtr=Number((growMemViews(),HEAPU64)[(entryPtr+40)/8]);var textureViewPtr=Number((growMemViews(),HEAPU64)[(entryPtr+48)/8]);var externalTexturePtr=0;WebGPU.iterateExtensions(entryPtr,{327681:ptr=>{externalTexturePtr=Number((growMemViews(),HEAPU64)[(ptr+16)/8])}});var resource;if(bufferPtr){var size=readI53FromI64(entryPtr+32);if(size==-1)size=undefined;resource={buffer:WebGPU.getJsObject(bufferPtr),offset:readI53FromI64(entryPtr+24),size}}else{resource=WebGPU.getJsObject(samplerPtr||textureViewPtr||externalTexturePtr)}return{binding:(growMemViews(),HEAPU32)[(entryPtr+8)/4],resource}}function makeEntries(count,entriesPtrs){var entries=[];for(var i=0;i<count;++i){entries.push(makeEntry(entriesPtrs+56*i))}return entries}var desc={label:WebGPU.makeStringFromOptionalStringView(descriptor+8),layout:WebGPU.getJsObject(Number((growMemViews(),HEAPU64)[(descriptor+24)/8])),entries:makeEntries(Number((growMemViews(),HEAPU64)[(descriptor+32)/8]),Number((growMemViews(),HEAPU64)[(descriptor+40)/8]))};var device=WebGPU.getJsObject(devicePtr);var ptr=_emwgpuCreateBindGroup(0n);WebGPU.Internals.jsObjectInsert(ptr,device.createBindGroup(desc));return ptr})();return BigInt(ret)};var _wgpuDeviceCreateCommandEncoder=function(devicePtr,descriptor){devicePtr=bigintToI53Checked(devicePtr);descriptor=bigintToI53Checked(descriptor);var ret=(()=>{var desc;if(descriptor){desc={label:WebGPU.makeStringFromOptionalStringView(descriptor+8)}}var device=WebGPU.getJsObject(devicePtr);var ptr=_emwgpuCreateCommandEncoder(0n);WebGPU.Internals.jsObjectInsert(ptr,device.createCommandEncoder(desc));return ptr})();return BigInt(ret)};var _wgpuDeviceCreateComputePipeline=function(devicePtr,descriptor){devicePtr=bigintToI53Checked(devicePtr);descriptor=bigintToI53Checked(descriptor);var ret=(()=>{var desc=WebGPU.makeComputePipelineDesc(descriptor);var device=WebGPU.getJsObject(devicePtr);var ptr=_emwgpuCreateComputePipeline(0n);WebGPU.Internals.jsObjectInsert(ptr,device.createComputePipeline(desc));return ptr})();return BigInt(ret)};function _wgpuInstanceHasWGSLLanguageFeature(instance,featureEnumValue){instance=bigintToI53Checked(instance);if(!("wgslLanguageFeatures"in navigator.gpu)){return false}return navigator.gpu.wgslLanguageFeatures.has(WebGPU.WGSLLanguageFeatureName[featureEnumValue])}var _wgpuQueueSubmit=function(queuePtr,commandCount,commands){queuePtr=bigintToI53Checked(queuePtr);commandCount=bigintToI53Checked(commandCount);commands=bigintToI53Checked(commands);var queue=WebGPU.getJsObject(queuePtr);var cmds=Array.from((growMemViews(),HEAP64).subarray(commands/8,(commands+commandCount*8)/8),id=>WebGPU.getJsObject(id));queue.submit(cmds)};function _wgpuQueueWriteBuffer(queuePtr,bufferPtr,bufferOffset,data,size){queuePtr=bigintToI53Checked(queuePtr);bufferPtr=bigintToI53Checked(bufferPtr);bufferOffset=bigintToI53Checked(bufferOffset);data=bigintToI53Checked(data);size=bigintToI53Checked(size);var queue=WebGPU.getJsObject(queuePtr);var buffer=WebGPU.getJsObject(bufferPtr);var subarray=(growMemViews(),HEAPU8).subarray(data,data+size);queue.writeBuffer(buffer,bufferOffset,subarray,0,size)}var Asyncify={instrumentWasmImports(imports){var importPattern=/^(invoke_.*|__asyncjs__.*)$/;for(let[x,original]of Object.entries(imports)){if(typeof original=="function"){let isAsyncifyImport=original.isAsync||importPattern.test(x);if(isAsyncifyImport){imports[x]=original=new WebAssembly.Suspending(original)}}}},instrumentFunction(original){var wrapper=(...args)=>original(...args);return wrapper},instrumentWasmExports(exports){var exportPattern=/^(wllama_start|wllama_action|main|__main_argc_argv)$/;Asyncify.asyncExports=new Set;var ret={};for(let[x,original]of Object.entries(exports)){if(typeof original=="function"){let isAsyncifyExport=exportPattern.test(x);if(isAsyncifyExport){Asyncify.asyncExports.add(original);original=Asyncify.makeAsyncFunction(original)}var wrapper=Asyncify.instrumentFunction(original);ret[x]=wrapper}else{ret[x]=original}}return ret},asyncExports:null,isAsyncExport(func){return Asyncify.asyncExports?.has(func)},handleAsync:async startAsync=>{runtimeKeepalivePush();try{return await startAsync()}finally{runtimeKeepalivePop()}},handleSleep:startAsync=>Asyncify.handleAsync(()=>new Promise(startAsync)),makeAsyncFunction(original){return WebAssembly.promising(original)}};var getCFunc=ident=>{var func=Module["_"+ident];return func};var writeArrayToMemory=(array,buffer)=>{(growMemViews(),HEAP8).set(array,buffer)};var ccall=(ident,returnType,argTypes,args,opts)=>{var toC={pointer:p=>BigInt(p),string:str=>{var ret=0;if(str!==null&&str!==undefined&&str!==0){ret=stringToUTF8OnStack(str)}return BigInt(ret)},array:arr=>{var ret=stackAlloc(arr.length);writeArrayToMemory(arr,ret);return BigInt(ret)}};function convertReturnValue(ret){if(returnType==="string"){return UTF8ToString(Number(ret))}if(returnType==="pointer")return Number(ret);if(returnType==="boolean")return Boolean(ret);return ret}var func=getCFunc(ident);var cArgs=[];var stack=0;if(args){for(var i=0;i<args.length;i++){var converter=toC[argTypes[i]];if(converter){if(stack===0)stack=stackSave();cArgs[i]=converter(args[i])}else{cArgs[i]=args[i]}}}var ret=func(...cArgs);function onDone(ret){if(stack!==0)stackRestore(stack);return convertReturnValue(ret)}var asyncMode=opts?.async;if(asyncMode)return ret.then(onDone);ret=onDone(ret);return ret};var cwrap=(ident,returnType,argTypes,opts)=>{var numericArgs=!argTypes||argTypes.every(type=>type==="number"||type==="boolean");var numericRet=returnType!=="string";if(numericRet&&numericArgs&&!opts){return getCFunc(ident)}return(...args)=>ccall(ident,returnType,argTypes,args,opts)};var FS_createPath=(...args)=>FS.createPath(...args);var FS_unlink=(...args)=>FS.unlink(...args);var FS_createLazyFile=(...args)=>FS.createLazyFile(...args);var FS_createDevice=(...args)=>FS.createDevice(...args);PThread.init();FS.createPreloadedFile=FS_createPreloadedFile;FS.preloadFile=FS_preloadFile;FS.staticInit();{initMemory();if(Module["noExitRuntime"])noExitRuntime=Module["noExitRuntime"];if(Module["preloadPlugins"])preloadPlugins=Module["preloadPlugins"];if(Module["print"])out=Module["print"];if(Module["printErr"])err=Module["printErr"];if(Module["wasmBinary"])wasmBinary=Module["wasmBinary"];if(Module["arguments"])arguments_=Module["arguments"];if(Module["thisProgram"])thisProgram=Module["thisProgram"];if(Module["preInit"]){if(typeof Module["preInit"]=="function")Module["preInit"]=[Module["preInit"]];while(Module["preInit"].length>0){Module["preInit"].shift()()}}}Module["ENV"]=ENV;Module["mmapAlloc"]=mmapAlloc;Module["wasmMemory"]=wasmMemory;Module["addRunDependency"]=addRunDependency;Module["removeRunDependency"]=removeRunDependency;Module["ccall"]=ccall;Module["cwrap"]=cwrap;Module["FS_preloadFile"]=FS_preloadFile;Module["FS_unlink"]=FS_unlink;Module["FS_createPath"]=FS_createPath;Module["FS_createDevice"]=FS_createDevice;Module["FS"]=FS;Module["FS_createDataFile"]=FS_createDataFile;Module["FS_createLazyFile"]=FS_createLazyFile;Module["MEMFS"]=MEMFS;var proxiedFunctionTable=[_proc_exit,exitOnMainThread,pthreadCreateProxied,___syscall_fcntl64,___syscall_fstat64,___syscall_getcwd,___syscall_getdents64,___syscall_ioctl,___syscall_lstat64,___syscall_newfstatat,___syscall_openat,___syscall_stat64,__mmap_js,__munmap_js,_environ_get,_environ_sizes_get,_fd_close,_fd_read,_fd_seek,_fd_write];function __asyncjs__js_file_read(path_ptr,offset,req_size,out_ptr){return Asyncify.handleAsync(async()=>await _wllama_js_file_read(UTF8ToString(Number(path_ptr)),Number(offset),Number(req_size),Number(out_ptr)))}__asyncjs__js_file_read.sig="jjjjj";var _malloc,_free,_wllama_malloc,_wllama_start,_wllama_action,_wllama_exit,_wllama_debug,_main,_emwgpuCreateBindGroup,_emwgpuCreateBindGroupLayout,_emwgpuCreateCommandBuffer,_emwgpuCreateCommandEncoder,_emwgpuCreateComputePassEncoder,_emwgpuCreateComputePipeline,_emwgpuCreateExternalTexture,_emwgpuCreatePipelineLayout,_emwgpuCreateQuerySet,_emwgpuCreateRenderBundle,_emwgpuCreateRenderBundleEncoder,_emwgpuCreateRenderPassEncoder,_emwgpuCreateRenderPipeline,_emwgpuCreateSampler,_emwgpuCreateSurface,_emwgpuCreateTexture,_emwgpuCreateTextureView,_emwgpuCreateAdapter,_emwgpuCreateBuffer,_emwgpuCreateDevice,_emwgpuCreateQueue,_emwgpuCreateShaderModule,_emwgpuOnDeviceLostCompleted,_emwgpuOnMapAsyncCompleted,_emwgpuOnRequestAdapterCompleted,_emwgpuOnRequestDeviceCompleted,_emwgpuOnWorkDoneCompleted,_emwgpuOnUncapturedError,__emscripten_tls_init,_pthread_self,_emscripten_builtin_memalign,__emscripten_thread_init,__emscripten_thread_crashed,__emscripten_run_js_on_main_thread,__emscripten_thread_free_data,__emscripten_thread_exit,__emscripten_check_mailbox,_memalign,___trap,_emscripten_stack_set_limits,__emscripten_stack_restore,__emscripten_stack_alloc,_emscripten_stack_get_current,__indirect_function_table,wasmTable;function assignWasmExports(wasmExports){_malloc=wasmExports["malloc"];_free=wasmExports["free"];_wllama_malloc=Module["_wllama_malloc"]=wasmExports["wllama_malloc"];_wllama_start=Module["_wllama_start"]=wasmExports["wllama_start"];_wllama_action=Module["_wllama_action"]=wasmExports["wllama_action"];_wllama_exit=Module["_wllama_exit"]=wasmExports["wllama_exit"];_wllama_debug=Module["_wllama_debug"]=wasmExports["wllama_debug"];_main=Module["_main"]=wasmExports["main"];_emwgpuCreateBindGroup=wasmExports["emwgpuCreateBindGroup"];_emwgpuCreateBindGroupLayout=wasmExports["emwgpuCreateBindGroupLayout"];_emwgpuCreateCommandBuffer=wasmExports["emwgpuCreateCommandBuffer"];_emwgpuCreateCommandEncoder=wasmExports["emwgpuCreateCommandEncoder"];_emwgpuCreateComputePassEncoder=wasmExports["emwgpuCreateComputePassEncoder"];_emwgpuCreateComputePipeline=wasmExports["emwgpuCreateComputePipeline"];_emwgpuCreateExternalTexture=wasmExports["emwgpuCreateExternalTexture"];_emwgpuCreatePipelineLayout=wasmExports["emwgpuCreatePipelineLayout"];_emwgpuCreateQuerySet=wasmExports["emwgpuCreateQuerySet"];_emwgpuCreateRenderBundle=wasmExports["emwgpuCreateRenderBundle"];_emwgpuCreateRenderBundleEncoder=wasmExports["emwgpuCreateRenderBundleEncoder"];_emwgpuCreateRenderPassEncoder=wasmExports["emwgpuCreateRenderPassEncoder"];_emwgpuCreateRenderPipeline=wasmExports["emwgpuCreateRenderPipeline"];_emwgpuCreateSampler=wasmExports["emwgpuCreateSampler"];_emwgpuCreateSurface=wasmExports["emwgpuCreateSurface"];_emwgpuCreateTexture=wasmExports["emwgpuCreateTexture"];_emwgpuCreateTextureView=wasmExports["emwgpuCreateTextureView"];_emwgpuCreateAdapter=wasmExports["emwgpuCreateAdapter"];_emwgpuCreateBuffer=wasmExports["emwgpuCreateBuffer"];_emwgpuCreateDevice=wasmExports["emwgpuCreateDevice"];_emwgpuCreateQueue=wasmExports["emwgpuCreateQueue"];_emwgpuCreateShaderModule=wasmExports["emwgpuCreateShaderModule"];_emwgpuOnDeviceLostCompleted=wasmExports["emwgpuOnDeviceLostCompleted"];_emwgpuOnMapAsyncCompleted=wasmExports["emwgpuOnMapAsyncCompleted"];_emwgpuOnRequestAdapterCompleted=wasmExports["emwgpuOnRequestAdapterCompleted"];_emwgpuOnRequestDeviceCompleted=wasmExports["emwgpuOnRequestDeviceCompleted"];_emwgpuOnWorkDoneCompleted=wasmExports["emwgpuOnWorkDoneCompleted"];_emwgpuOnUncapturedError=wasmExports["emwgpuOnUncapturedError"];__emscripten_tls_init=wasmExports["_emscripten_tls_init"];_pthread_self=wasmExports["pthread_self"];_emscripten_builtin_memalign=wasmExports["emscripten_builtin_memalign"];__emscripten_thread_init=wasmExports["_emscripten_thread_init"];__emscripten_thread_crashed=wasmExports["_emscripten_thread_crashed"];__emscripten_run_js_on_main_thread=wasmExports["_emscripten_run_js_on_main_thread"];__emscripten_thread_free_data=wasmExports["_emscripten_thread_free_data"];__emscripten_thread_exit=wasmExports["_emscripten_thread_exit"];__emscripten_check_mailbox=wasmExports["_emscripten_check_mailbox"];_memalign=wasmExports["memalign"];___trap=wasmExports["__trap"];_emscripten_stack_set_limits=wasmExports["emscripten_stack_set_limits"];__emscripten_stack_restore=wasmExports["_emscripten_stack_restore"];__emscripten_stack_alloc=wasmExports["_emscripten_stack_alloc"];_emscripten_stack_get_current=wasmExports["emscripten_stack_get_current"];__indirect_function_table=wasmTable=wasmExports["__indirect_function_table"]}var wasmImports;function assignWasmImports(){wasmImports={__asyncjs__js_file_read,__pthread_create_js:___pthread_create_js,__syscall_fcntl64:___syscall_fcntl64,__syscall_getcwd:___syscall_getcwd,__syscall_getdents64:___syscall_getdents64,__syscall_ioctl:___syscall_ioctl,__syscall_openat:___syscall_openat,__syscall_stat64:___syscall_stat64,_abort_js:__abort_js,_emscripten_init_main_thread_js:__emscripten_init_main_thread_js,_emscripten_notify_mailbox_postmessage:__emscripten_notify_mailbox_postmessage,_emscripten_receive_on_main_thread_js:__emscripten_receive_on_main_thread_js,_emscripten_thread_cleanup:__emscripten_thread_cleanup,_emscripten_thread_mailbox_await:__emscripten_thread_mailbox_await,_emscripten_thread_set_strongref:__emscripten_thread_set_strongref,_localtime_js:__localtime_js,_mmap_js:__mmap_js,_munmap_js:__munmap_js,_tzset_js:__tzset_js,clock_time_get:_clock_time_get,emscripten_check_blocking_allowed:_emscripten_check_blocking_allowed,emscripten_date_now:_emscripten_date_now,emscripten_exit_with_live_runtime:_emscripten_exit_with_live_runtime,emscripten_get_callstack:_emscripten_get_callstack,emscripten_get_heap_max:_emscripten_get_heap_max,emscripten_get_now:_emscripten_get_now,emscripten_has_asyncify:_emscripten_has_asyncify,emscripten_num_logical_cores:_emscripten_num_logical_cores,emscripten_resize_heap:_emscripten_resize_heap,emwgpuAdapterRequestDevice:_emwgpuAdapterRequestDevice,emwgpuBufferDestroy:_emwgpuBufferDestroy,emwgpuBufferGetConstMappedRange:_emwgpuBufferGetConstMappedRange,emwgpuBufferMapAsync:_emwgpuBufferMapAsync,emwgpuBufferUnmap:_emwgpuBufferUnmap,emwgpuDelete:_emwgpuDelete,emwgpuDeviceCreateBuffer:_emwgpuDeviceCreateBuffer,emwgpuDeviceCreateShaderModule:_emwgpuDeviceCreateShaderModule,emwgpuDeviceDestroy:_emwgpuDeviceDestroy,emwgpuInstanceRequestAdapter:_emwgpuInstanceRequestAdapter,emwgpuQueueOnSubmittedWorkDone:_emwgpuQueueOnSubmittedWorkDone,emwgpuWaitAny:_emwgpuWaitAny,environ_get:_environ_get,environ_sizes_get:_environ_sizes_get,exit:_exit,fd_close:_fd_close,fd_read:_fd_read,fd_seek:_fd_seek,fd_write:_fd_write,memory:wasmMemory,random_get:_random_get,wgpuAdapterGetInfo:_wgpuAdapterGetInfo,wgpuAdapterGetLimits:_wgpuAdapterGetLimits,wgpuAdapterHasFeature:_wgpuAdapterHasFeature,wgpuBufferGetSize:_wgpuBufferGetSize,wgpuCommandEncoderBeginComputePass:_wgpuCommandEncoderBeginComputePass,wgpuCommandEncoderCopyBufferToBuffer:_wgpuCommandEncoderCopyBufferToBuffer,wgpuCommandEncoderFinish:_wgpuCommandEncoderFinish,wgpuComputePassEncoderDispatchWorkgroups:_wgpuComputePassEncoderDispatchWorkgroups,wgpuComputePassEncoderEnd:_wgpuComputePassEncoderEnd,wgpuComputePassEncoderSetBindGroup:_wgpuComputePassEncoderSetBindGroup,wgpuComputePassEncoderSetPipeline:_wgpuComputePassEncoderSetPipeline,wgpuComputePipelineGetBindGroupLayout:_wgpuComputePipelineGetBindGroupLayout,wgpuDeviceCreateBindGroup:_wgpuDeviceCreateBindGroup,wgpuDeviceCreateCommandEncoder:_wgpuDeviceCreateCommandEncoder,wgpuDeviceCreateComputePipeline:_wgpuDeviceCreateComputePipeline,wgpuInstanceHasWGSLLanguageFeature:_wgpuInstanceHasWGSLLanguageFeature,wgpuQueueSubmit:_wgpuQueueSubmit,wgpuQueueWriteBuffer:_wgpuQueueWriteBuffer}}function applySignatureConversions(wasmExports){wasmExports=Object.assign({},wasmExports);var makeWrapper_pp=f=>a0=>Number(f(BigInt(a0)));var makeWrapper__p=f=>a0=>f(BigInt(a0));var makeWrapper___PP=f=>(a0,a1,a2)=>f(a0,BigInt(a1?a1:0),BigInt(a2?a2:0));var makeWrapper_p=f=>()=>Number(f());var makeWrapper_ppp=f=>(a0,a1)=>Number(f(BigInt(a0),BigInt(a1)));var makeWrapper__p_____=f=>(a0,a1,a2,a3,a4,a5)=>f(BigInt(a0),a1,a2,a3,a4,a5);var makeWrapper___p_p_=f=>(a0,a1,a2,a3,a4)=>f(a0,BigInt(a1),a2,BigInt(a3),a4);var makeWrapper__pp=f=>(a0,a1)=>f(BigInt(a0),BigInt(a1));wasmExports["malloc"]=makeWrapper_pp(wasmExports["malloc"]);wasmExports["free"]=makeWrapper__p(wasmExports["free"]);wasmExports["main"]=makeWrapper___PP(wasmExports["main"]);wasmExports["pthread_self"]=makeWrapper_p(wasmExports["pthread_self"]);wasmExports["emscripten_builtin_memalign"]=makeWrapper_ppp(wasmExports["emscripten_builtin_memalign"]);wasmExports["_emscripten_thread_init"]=makeWrapper__p_____(wasmExports["_emscripten_thread_init"]);wasmExports["_emscripten_run_js_on_main_thread"]=makeWrapper___p_p_(wasmExports["_emscripten_run_js_on_main_thread"]);wasmExports["_emscripten_thread_free_data"]=makeWrapper__p(wasmExports["_emscripten_thread_free_data"]);wasmExports["_emscripten_thread_exit"]=makeWrapper__p(wasmExports["_emscripten_thread_exit"]);wasmExports["memalign"]=makeWrapper_ppp(wasmExports["memalign"]);wasmExports["emscripten_stack_set_limits"]=makeWrapper__pp(wasmExports["emscripten_stack_set_limits"]);wasmExports["_emscripten_stack_restore"]=makeWrapper__p(wasmExports["_emscripten_stack_restore"]);wasmExports["_emscripten_stack_alloc"]=makeWrapper_pp(wasmExports["_emscripten_stack_alloc"]);wasmExports["emscripten_stack_get_current"]=makeWrapper_p(wasmExports["emscripten_stack_get_current"]);return wasmExports}async function callMain(){var entryFunction=_main;var argc=0;var argv=0;try{var ret=entryFunction(argc,BigInt(argv));ret=await ret;exitJS(ret,true);return ret}catch(e){return handleException(e)}}function run(){if(runDependencies>0){dependenciesFulfilled=run;return}if(ENVIRONMENT_IS_PTHREAD){initRuntime();return}preRun();if(runDependencies>0){dependenciesFulfilled=run;return}async function doRun(){Module["calledRun"]=true;if(ABORT)return;initRuntime();preMain();Module["onRuntimeInitialized"]?.();var noInitialRun=Module["noInitialRun"]||false;if(!noInitialRun)await callMain();postRun()}if(Module["setStatus"]){Module["setStatus"]("Running...");setTimeout(()=>{setTimeout(()=>Module["setStatus"](""),1);doRun()},1)}else{doRun()}}var wasmExports;if(!ENVIRONMENT_IS_PTHREAD){createWasm();run()}\n';

// src/worker.ts
var FILE_READ_REQ_EVENT = "fs.read_req";
var JSPI_STUB = `
if (!WebAssembly.Suspending) {
  // JSPI not available - stubs that keep the import/export tables valid.
  // Suspending wraps imports: identity is fine since async imports won't be called.
  WebAssembly.Suspending = function (fn) {
    // console.log(fn.toString());
    return fn;
  };
  // promising wraps exports: must return a Promise so ccall's ret.then() works.
  WebAssembly.promising = function (fn) {
    return function (...args) {
      try {
        return Promise.resolve(fn(...args));
      } catch (e) {
        return Promise.reject(e);
      }
    };
  };
}
`;
var ProxyToWorker = class {
  // filename -> Blob for async reads
  constructor(resources, nbThread, suppressNativeLog, logger) {
    __publicField(this, "resources");
    __publicField(this, "logger");
    __publicField(this, "suppressNativeLog");
    __publicField(this, "taskQueue", []);
    __publicField(this, "taskId", 1);
    __publicField(this, "resultQueue", []);
    __publicField(this, "busy", false);
    // is the work loop is running?
    __publicField(this, "worker");
    __publicField(this, "multiThread");
    __publicField(this, "nbThread");
    __publicField(this, "useAsyncFile");
    __publicField(this, "fileBlobs", /* @__PURE__ */ new Map());
    this.resources = resources;
    this.nbThread = nbThread;
    this.multiThread = nbThread > 0;
    this.logger = logger;
    this.suppressNativeLog = suppressNativeLog;
    this.useAsyncFile = canUseAsyncFileRead(resources.compat);
  }
  getModuleCode() {
    return __async(this, null, function* () {
      if (!this.resources.jsPath) {
        if (this.resources.compat) {
          throw new Error(
            "compat mode is enabled but no jsPath was provided. Pass a worker JS via setCompat() or install @wllama/wllama-compat."
          );
        }
        return WLLAMA_EMSCRIPTEN_CODE;
      } else if (this.resources.jsPath.code) {
        return this.resources.jsPath.code;
      } else if (isString(this.resources.jsPath)) {
        const response = yield fetch(this.resources.jsPath);
        if (!response.ok) {
          throw new Error(
            `Failed to fetch worker code from ${this.resources.jsPath}`
          );
        }
        return yield response.text();
      } else {
        throw new Error("No JS code provided for worker");
      }
    });
  }
  moduleInit(ggufFiles) {
    return __async(this, null, function* () {
      let moduleCode = JSPI_STUB + (yield this.getModuleCode());
      if (this.resources.noWebGPU) {
        moduleCode = 'try{Object.defineProperty(WorkerNavigator.prototype,"gpu",{get:()=>({requestAdapter:async()=>null})});}catch(e){}' + moduleCode;
      }
      let mainModuleCode = moduleCode.replace("var Module", "var ___Module");
      const runOptions = {
        pathConfig: {
          "wllama.wasm": this.resources.wasmPath
        },
        nbThread: this.nbThread,
        compat: this.resources.compat
      };
      const completeCode = [
        `const RUN_OPTIONS = ${JSON.stringify(runOptions)};`,
        `function wModuleInit() { ${mainModuleCode}; return Module; }`,
        LLAMA_CPP_WORKER_CODE
      ].join(";\n\n");
      this.worker = createWorker(completeCode);
      this.worker.onmessage = this.onRecvMsg.bind(this);
      this.worker.onerror = this.logger.error;
      const res = yield this.pushTask({
        verb: "module.init",
        args: [
          new Blob([moduleCode], { type: "text/javascript" }),
          this.useAsyncFile
        ],
        callbackId: this.taskId++
      });
      const nativeFiles = [];
      for (const file of ggufFiles) {
        const needAllocBuffer = !this.useAsyncFile;
        const id = yield this.fileAlloc(
          file.name,
          file.blob.size,
          needAllocBuffer
        );
        nativeFiles.push(__spreadValues({ id }, file));
        if (this.useAsyncFile) {
          this.fileBlobs.set(file.name, file.blob);
        }
      }
      if (!this.useAsyncFile) {
        yield Promise.all(
          nativeFiles.map((file) => {
            return this.fileWrite(file.id, file.blob);
          })
        );
      }
      return res;
    });
  }
  wllamaStart() {
    return __async(this, null, function* () {
      const result = yield this.pushTask({
        verb: "wllama.start",
        args: [],
        callbackId: this.taskId++
      });
      const parsedResult = this.parseResult(result);
      return parsedResult;
    });
  }
  wllamaAction(name, body) {
    return __async(this, null, function* () {
      const encodedMsg = glueSerialize(body);
      const result = yield this.pushTask({
        verb: "wllama.action",
        args: [name, encodedMsg],
        callbackId: this.taskId++
      });
      const parsedResult = glueDeserialize(result);
      return parsedResult;
    });
  }
  wllamaExit() {
    return __async(this, null, function* () {
      if (this.worker) {
        this.worker.terminate();
      }
    });
  }
  wllamaDebug() {
    return __async(this, null, function* () {
      const result = yield this.pushTask({
        verb: "wllama.debug",
        args: [],
        callbackId: this.taskId++
      });
      return JSON.parse(result);
    });
  }
  ///////////////////////////////////////
  /**
   * Allocate a new file in heapfs
   * @returns fileId, to be used by fileWrite()
   */
  fileAlloc(fileName, size, allocBuffer) {
    return __async(this, null, function* () {
      const result = yield this.pushTask({
        verb: "fs.alloc",
        args: [fileName, size, allocBuffer],
        callbackId: this.taskId++
      });
      return result.fileId;
    });
  }
  /**
   * Write a Blob to heapfs
   */
  fileWrite(fileId, blob) {
    return __async(this, null, function* () {
      const reader = blob.stream().getReader();
      let offset = 0;
      while (true) {
        const { done, value } = yield reader.read();
        if (done) break;
        const size = value.byteLength;
        yield this.pushTask(
          {
            verb: "fs.write",
            args: [fileId, value, offset],
            callbackId: this.taskId++
          },
          // @ts-ignore Type 'ArrayBufferLike' is not assignable to type 'ArrayBuffer'
          [value.buffer]
        );
        offset += size;
      }
    });
  }
  fileReadResponse(name, offset, size) {
    return __async(this, null, function* () {
      var _a;
      try {
        const blob = this.fileBlobs.get(name);
        if (!blob) {
          throw new Error(`blob not found for name="${name}"`);
        }
        const chunk = blob.slice(offset, offset + size);
        const buffer = yield chunk.arrayBuffer();
        this.worker.postMessage(
          { verb: "fs.read_res", args: [buffer] },
          { transfer: [buffer] }
        );
      } catch (err) {
        this.logger.error("fileReadResponse failed, terminating worker:", err);
        (_a = this.worker) == null ? void 0 : _a.terminate();
        this.worker = void 0;
        this.abort(`File read failed: ${err}`, err.stack || "");
      }
    });
  }
  /**
   * Parse JSON result returned by cpp code.
   * Throw new Error if "__exception" is present in the response
   *
   * TODO: get rid of this function once everything is migrated to Glue
   */
  parseResult(result) {
    const parsedResult = JSON.parse(result);
    if (parsedResult && parsedResult["error"]) {
      throw new WllamaRuntimeError("Unknown error, please see console.log", "");
    }
    return parsedResult;
  }
  /**
   * Push a new task to taskQueue
   */
  pushTask(param, buffers) {
    return new Promise((resolve, reject) => {
      this.taskQueue.push({ resolve, reject, param, buffers });
      this.runTaskLoop();
    });
  }
  /**
   * Main loop for processing tasks
   */
  runTaskLoop() {
    return __async(this, null, function* () {
      var _a;
      if (this.busy) {
        return;
      }
      this.busy = true;
      while (true) {
        const task = this.taskQueue.shift();
        if (!task) break;
        this.resultQueue.push(task);
        this.worker.postMessage(
          task.param,
          isSafariMobile() ? void 0 : {
            transfer: (_a = task.buffers) != null ? _a : []
          }
        );
      }
      this.busy = false;
    });
  }
  /**
   * Handle messages from worker
   */
  onRecvMsg(e) {
    if (!e.data) return;
    const { verb, args } = e.data;
    const isCompatBuild = this.resources.compat;
    if (verb && verb.startsWith("console.")) {
      if (this.suppressNativeLog) {
        return;
      }
      if (verb.endsWith("debug")) this.logger.debug(...args);
      if (verb.endsWith("log")) this.logger.log(...args);
      if (verb.endsWith("warn")) this.logger.warn(...args);
      if (verb.endsWith("error")) this.logger.error(...args);
      return;
    } else if (verb === "signal.abort") {
      const [signalType, message, rawStack, originalErr] = args;
      if (originalErr) {
        this.logger.error(originalErr);
      }
      (() => __async(this, null, function* () {
        let stack = "";
        let newMsg = message.replace(
          "Build with -sASSERTIONS for more info.",
          ""
        );
        if (signalType === "abort") {
          newMsg = `(ABORT) ${newMsg}`;
          stack = rawStack.replace(/\|/g, "\n");
        } else if (signalType === "exception") {
          stack = rawStack;
        }
        const decoded = yield Debug.decodeStackTrace(stack, isCompatBuild);
        this.logger.error(`Stack trace (${signalType}):
` + decoded);
        this.abort(newMsg, decoded);
      }))();
      return;
    }
    if (verb === FILE_READ_REQ_EVENT) {
      const [name, offset, size] = args;
      this.fileReadResponse(name, offset, size).catch(() => {
      });
      return;
    }
    const { callbackId, result, err } = e.data;
    if (callbackId) {
      const idx = this.resultQueue.findIndex(
        (t) => t.param.callbackId === callbackId
      );
      if (idx !== -1) {
        const waitingTask = this.resultQueue.splice(idx, 1)[0];
        if (err) waitingTask.reject(err);
        else waitingTask.resolve(result);
      } else {
        this.logger.error(
          `Cannot find waiting task with callbackId = ${callbackId}`
        );
      }
    }
  }
  abort(text, stack) {
    const error = new WllamaRuntimeError(
      text.length == 0 ? "(unknown error)" : text,
      stack
    );
    while (this.resultQueue.length > 0) {
      const waitingTask = this.resultQueue.pop();
      if (!waitingTask) break;
      waitingTask.reject(error);
    }
    while (this.taskQueue.length > 0) {
      const pendingTask = this.taskQueue.pop();
      if (!pendingTask) break;
      pendingTask.reject(error);
    }
  }
};

// src/huggingface.ts
var HF_BASE = "https://huggingface.co";
var DEFAULT_QUANTS = ["Q4_K_M", "Q8_0"];
function fetchRepoFiles(repo, token) {
  return __async(this, null, function* () {
    var _a;
    const url = `${HF_BASE}/api/models/${repo}/tree/main?recursive=true`;
    const headers = { Accept: "application/json" };
    if (token) {
      headers["Authorization"] = `Bearer ${token}`;
    }
    const res = yield fetch(url, { headers });
    if (!res.ok) {
      let msg = res.statusText;
      try {
        msg = (_a = (yield res.json()).error) != null ? _a : msg;
      } catch (e) {
      }
      throw new Error(`HF API error (${res.status}): ${msg}`);
    }
    return res.json();
  });
}
function firstShardPath(files, path) {
  const m = path.match(/^(.+)-(\d{5})-of-(\d{5})\.gguf$/i);
  if (!m) return path;
  const first = `${m[1]}-00001-of-${m[3]}.gguf`;
  return files.some((f) => f.path === first) ? first : path;
}
function selectFile(files, quant, mmprojOnly) {
  const candidates = files.filter((f) => {
    if (f.type !== "file" || !f.path.toLowerCase().endsWith(".gguf"))
      return false;
    const ismmproj = f.path.toLowerCase().includes("mmproj");
    return mmprojOnly ? ismmproj : !ismmproj;
  });
  if (candidates.length === 0) return null;
  if (quant) {
    const upper = quant.toUpperCase();
    const match = candidates.find((f) => f.path.toUpperCase().includes(upper));
    if (match) return firstShardPath(candidates, match.path);
    return null;
  }
  for (const q of DEFAULT_QUANTS) {
    const match = candidates.find((f) => f.path.toUpperCase().includes(q));
    if (match) return firstShardPath(candidates, match.path);
  }
  return firstShardPath(candidates, candidates[0].path);
}
function getHFModelSource(config) {
  return __async(this, null, function* () {
    const { repo, file, quant, mmprojFile, mmprojQuant, hfToken } = config;
    const files = yield fetchRepoFiles(repo, hfToken);
    const modelPath = file != null ? file : selectFile(files, quant, false);
    if (!modelPath) {
      throw new Error(`No GGUF file found in repo "${repo}"`);
    }
    const source = {
      url: `${HF_BASE}/${repo}/resolve/main/${modelPath}`
    };
    if (mmprojFile || mmprojQuant !== void 0) {
      const mmpath = mmprojFile != null ? mmprojFile : selectFile(files, mmprojQuant, true);
      if (mmpath) {
        source.mmprojUrl = `${HF_BASE}/${repo}/resolve/main/${mmpath}`;
      }
    }
    if (hfToken) {
      const params = new URLSearchParams({ token: hfToken });
      source.url += `?${params}`;
      if (source.mmprojUrl) {
        source.mmprojUrl += `?${params}`;
      }
    }
    return source;
  });
}
function getHFFileSHA256(url, headers) {
  return __async(this, null, function* () {
    if (!url.includes("/resolve/")) return void 0;
    const rawUrl = url.replace("/resolve/", "/raw/");
    try {
      const text = yield fetch(rawUrl, { headers }).then((r) => r.text());
      const match = text.match(/^oid sha256:([0-9a-f]{64})$/m);
      return match ? match[1] : void 0;
    } catch (e) {
      return void 0;
    }
  });
}

// src/storage/opfs.ts
var OPFSBackend = class {
  isSupported() {
    var _a;
    return typeof navigator !== "undefined" && "storage" in navigator && !!((_a = navigator.storage) == null ? void 0 : _a.getDirectory);
  }
  read(key) {
    return __async(this, null, function* () {
      try {
        const cacheDir = yield getCacheDir();
        const fileHandle = yield cacheDir.getFileHandle(key);
        return yield fileHandle.getFile();
      } catch (e) {
        return null;
      }
    });
  }
  write(key, stream) {
    return __async(this, null, function* () {
      const writable = yield openWritable(key);
      yield writable.truncate(0);
      const reader = stream.getReader();
      try {
        while (true) {
          const { done, value } = yield reader.read();
          if (done) break;
          yield writable.write(value);
        }
      } finally {
        yield writable.close();
      }
    });
  }
  getSize(key) {
    return __async(this, null, function* () {
      try {
        const cacheDir = yield getCacheDir();
        const fileHandle = yield cacheDir.getFileHandle(key);
        const file = yield fileHandle.getFile();
        return file.size;
      } catch (e) {
        return -1;
      }
    });
  }
  list() {
    return __async(this, null, function* () {
      const cacheDir = yield getCacheDir();
      const result = [];
      try {
        for (var iter = __forAwait(cacheDir.entries()), more, temp, error; more = !(temp = yield iter.next()).done; more = false) {
          const [name, handle] = temp.value;
          if (handle.kind === "file") {
            const file = yield handle.getFile();
            result.push({ key: name, size: file.size });
          }
        }
      } catch (temp) {
        error = [temp];
      } finally {
        try {
          more && (temp = iter.return) && (yield temp.call(iter));
        } finally {
          if (error)
            throw error[0];
        }
      }
      return result;
    });
  }
  delete(key) {
    return __async(this, null, function* () {
      try {
        const cacheDir = yield getCacheDir();
        yield cacheDir.removeEntry(key);
      } catch (e) {
        if ((e == null ? void 0 : e.name) !== "NotFoundError") throw e;
      }
    });
  }
};
function getCacheDir() {
  return __async(this, null, function* () {
    const opfsRoot = yield navigator.storage.getDirectory();
    return opfsRoot.getDirectoryHandle("cache", { create: true });
  });
}
function openWritable(fileName) {
  return __async(this, null, function* () {
    const worker = createWorker(OPFS_UTILS_WORKER_CODE);
    let pResolve;
    let pReject;
    worker.onmessage = (e) => {
      if (e.data.ok) pResolve(null);
      else if (e.data.err) pReject(e.data.err);
    };
    worker.onerror = (e) => {
      var _a;
      return pReject == null ? void 0 : pReject((_a = e.message) != null ? _a : e);
    };
    const workerExec = (data) => new Promise((resolve, reject) => {
      pResolve = resolve;
      pReject = reject;
      worker.postMessage(
        data,
        isSafariMobile() ? void 0 : { transfer: "buf" in data && data.buf ? [data.buf.buffer] : [] }
      );
    });
    yield workerExec({ action: "open", filename: fileName });
    return {
      truncate: () => __async(this, null, function* () {
      }),
      write: (value) => workerExec({ action: "write", buf: value }),
      close: () => __async(this, null, function* () {
        yield workerExec({ action: "close" });
        worker.terminate();
      })
    };
  });
}

// src/storage/cos.ts
function makeHash(key) {
  return { algorithm: "SHA-256", value: key };
}
var COSInternalBackend = class {
  isSupported() {
    return typeof navigator !== "undefined" && "crossOriginStorage" in navigator;
  }
  // IMPORTANT: key must be SHA-256 hash of the data
  read(key) {
    return __async(this, null, function* () {
      try {
        const handle = yield navigator.crossOriginStorage.requestFileHandle(
          makeHash(key)
        );
        return handle.getFile();
      } catch (e) {
        return null;
      }
    });
  }
  // IMPORTANT: key must be SHA-256 hash of the data
  write(key, stream) {
    return __async(this, null, function* () {
      const handle = yield navigator.crossOriginStorage.requestFileHandle(
        makeHash(key),
        { create: true }
      );
      const writable = yield handle.createWritable();
      const reader = stream.getReader();
      try {
        while (true) {
          const { done, value } = yield reader.read();
          if (done) break;
          yield writable.write(value);
        }
      } finally {
        yield writable.close();
      }
    });
  }
  // IMPORTANT: key must be SHA-256 hash of the data
  getSize(key) {
    return __async(this, null, function* () {
      try {
        const handle = yield navigator.crossOriginStorage.requestFileHandle(
          makeHash(key)
        );
        const file = yield handle.getFile();
        return file.size;
      } catch (e) {
        return -1;
      }
    });
  }
  list() {
    return __async(this, null, function* () {
      throw new Error("not implemented");
    });
  }
  delete(_key) {
    return __async(this, null, function* () {
      throw new Error("not implemented");
    });
  }
};
var COSBackend = class {
  constructor() {
    __publicField(this, "cos", new COSInternalBackend());
    __publicField(this, "priv", new OPFSBackend());
  }
  isSupported() {
    return this.priv.isSupported();
  }
  read(key, hint) {
    return __async(this, null, function* () {
      if ((hint == null ? void 0 : hint.sha256) && this.cos.isSupported()) {
        const blob = yield this.cos.read(hint.sha256);
        if (blob) return blob;
      }
      return this.priv.read(key);
    });
  }
  write(key, stream, hint) {
    return __async(this, null, function* () {
      if ((hint == null ? void 0 : hint.sha256) && this.cos.isSupported()) {
        yield this.cos.write(hint.sha256, stream);
      } else {
        yield this.priv.write(key, stream);
      }
    });
  }
  getSize(key, hint) {
    return __async(this, null, function* () {
      if ((hint == null ? void 0 : hint.sha256) && this.cos.isSupported()) {
        const size = yield this.cos.getSize(hint.sha256);
        if (size !== -1) return size;
      }
      return this.priv.getSize(key);
    });
  }
  list() {
    return __async(this, null, function* () {
      return this.priv.list();
    });
  }
  delete(key) {
    return __async(this, null, function* () {
      return this.priv.delete(key);
    });
  }
};

// src/cache-manager.ts
var PREFIX_METADATA = "__metadata__";
var POLYFILL_ETAG = "polyfill_for_older_version";
function hintFromMetadata(metadata) {
  if (!metadata) return void 0;
  if (metadata.sha256) return { sha256: metadata.sha256 };
  return void 0;
}
var CacheManager = class {
  /**
   * @param backends Array of storage backends to use, in order of preference ; if first is available, use it, otherwise try the next one.
   */
  constructor(backends = [new COSBackend()]) {
    __publicField(this, "sb");
    for (const backend of backends) {
      if (backend.isSupported()) {
        this.sb = backend;
        return;
      }
    }
    throw new Error("No supported storage backend found");
  }
  /**
   * Convert a given URL into a storage key.
   *
   * Format: `${hashSHA1(fullURL)}_${fileName}`
   */
  getNameFromURL(url) {
    return __async(this, null, function* () {
      return urlToFileName(url, "");
    });
  }
  /**
   * @deprecated Use `download()` instead
   *
   * Write a new file to cache. This will overwrite existing file.
   *
   * @param name The file name returned by `getNameFromURL()` or `list()`
   */
  write(name, stream, metadata) {
    return __async(this, null, function* () {
      yield this.sb.write(name, stream);
      yield this.writeMetadata(name, metadata);
    });
  }
  download(_0) {
    return __async(this, arguments, function* (url, options = {}) {
      var _a, _b, _c, _d;
      const fileKey = yield urlToFileName(url, "");
      const sha256 = yield getHFFileSHA256(url, (_a = options.headers) != null ? _a : {});
      const hint = sha256 ? { sha256 } : void 0;
      const cachedSize = yield this.sb.getSize(fileKey, hint);
      if (cachedSize !== -1) {
        const metadata2 = yield this.readMetadata(fileKey);
        if ((metadata2 == null ? void 0 : metadata2.originalURL) === url && metadata2.originalSize === cachedSize) {
          return;
        }
        const head = yield fetch(url, __spreadValues({
          method: "HEAD"
        }, options.headers ? { headers: options.headers } : {}));
        const originalSize = parseInt(
          (_b = head.headers.get("content-length")) != null ? _b : "0",
          10
        );
        const etag2 = (head.headers.get("etag") || "").replace(
          /[^A-Za-z0-9]/g,
          ""
        );
        if (originalSize > 0 && originalSize === cachedSize) {
          yield this.writeMetadata(fileKey, __spreadValues({
            originalURL: url,
            originalSize,
            etag: etag2,
            sha256
          }, (_c = options.metadataAdditional) != null ? _c : {}));
          return;
        }
        yield this.sb.delete(fileKey);
        yield this.sb.delete(`${PREFIX_METADATA}${fileKey}`);
      }
      const response = yield fetch(url, __spreadValues(__spreadValues({}, options.headers ? { headers: options.headers } : {}), options.signal ? { signal: options.signal } : {}));
      if (!response.ok || !response.body) {
        throw new Error(`Failed to fetch ${url}: HTTP ${response.status}`);
      }
      const contentLength = response.headers.get("content-length");
      const etag = (response.headers.get("etag") || "").replace(
        /[^A-Za-z0-9]/g,
        ""
      );
      const total = parseInt(contentLength != null ? contentLength : "0", 10);
      const progressCallback = options.progressCallback;
      let loaded = 0;
      let lastProgressAt = 0;
      const progressStream = new TransformStream({
        transform(chunk, controller) {
          loaded += chunk.byteLength;
          if (progressCallback) {
            const now = Date.now();
            if (now - lastProgressAt > 100) {
              lastProgressAt = now;
              progressCallback({ loaded, total });
            }
          }
          controller.enqueue(chunk);
        },
        flush() {
          progressCallback == null ? void 0 : progressCallback({ loaded, total: total || loaded });
        }
      });
      const metadata = __spreadValues({
        originalURL: url,
        originalSize: total,
        etag
      }, (_d = options.metadataAdditional) != null ? _d : {});
      if (sha256) {
        metadata.sha256 = sha256;
      }
      yield this.sb.write(
        fileKey,
        response.body.pipeThrough(progressStream),
        hint
      );
      yield this.writeMetadata(fileKey, metadata);
    });
  }
  /**
   * Open a file in cache for reading
   *
   * @param nameOrURL The file name returned by `getNameFromURL()` or `list()`, or the original URL of the remote file
   * @returns Blob, or null if file does not exist
   */
  open(nameOrURL) {
    return __async(this, null, function* () {
      const hint1 = hintFromMetadata(yield this.getMetadata(nameOrURL));
      const direct = yield this.sb.read(nameOrURL, hint1);
      if (direct) return direct;
      const key = yield urlToFileName(nameOrURL, "");
      const hint2 = hintFromMetadata(yield this.getMetadata(key));
      return this.sb.read(key, hint2);
    });
  }
  /**
   * Get the size of a file in stored cache
   *
   * NOTE: in case the download is stopped mid-way (i.e. user close browser tab), the file maybe corrupted, size maybe different from `metadata.originalSize`
   *
   * @param name The file name returned by `getNameFromURL()` or `list()`
   * @returns number of bytes, or -1 if file does not exist
   */
  getSize(name) {
    return __async(this, null, function* () {
      const hint = hintFromMetadata(yield this.getMetadata(name));
      return this.sb.getSize(name, hint);
    });
  }
  /**
   * Get metadata of a cached file
   */
  getMetadata(name) {
    return __async(this, null, function* () {
      const metadata = yield this.readMetadata(name);
      if (metadata) return metadata;
      const cachedSize = yield this.sb.getSize(name);
      return cachedSize > 0 ? (
        // files created by older version of wllama don't have metadata; polyfill it
        {
          etag: POLYFILL_ETAG,
          originalSize: cachedSize,
          originalURL: ""
        }
      ) : (
        // cached file not found
        null
      );
    });
  }
  /**
   * Same as `getMetadata()`, but without polyfill. Returns null if the file has no metadata.
   */
  readMetadata(name) {
    return __async(this, null, function* () {
      const blob = yield this.sb.read(`${PREFIX_METADATA}${name}`);
      if (!blob) return null;
      try {
        return yield new Response(blob).json();
      } catch (e) {
        return null;
      }
    });
  }
  /**
   * List all files currently in cache
   */
  list() {
    return __async(this, null, function* () {
      const all = yield this.sb.list();
      const metadataMap = {};
      for (const { key } of all) {
        if (key.startsWith(PREFIX_METADATA)) {
          const blob = yield this.sb.read(key);
          if (blob) {
            const meta = yield new Response(blob).json().catch(() => null);
            metadataMap[key.slice(PREFIX_METADATA.length)] = meta;
          }
        }
      }
      const result = [];
      for (const { key, size } of all) {
        if (!key.startsWith(PREFIX_METADATA)) {
          result.push({
            name: key,
            size,
            metadata: metadataMap[key] || {
              originalSize: size,
              originalURL: "",
              etag: ""
            }
          });
        }
      }
      return result;
    });
  }
  /**
   * Clear all files currently in cache
   */
  clear() {
    return __async(this, null, function* () {
      yield this.deleteMany(() => true);
    });
  }
  /**
   * Delete a single file in cache
   *
   * @param nameOrURL Can be either an URL or a name returned by `getNameFromURL()` or `list()`
   */
  delete(nameOrURL) {
    return __async(this, null, function* () {
      const name2 = yield this.getNameFromURL(nameOrURL);
      yield this.deleteMany(
        (entry) => entry.name === nameOrURL || entry.name === name2
      );
    });
  }
  /**
   * Delete multiple files in cache.
   *
   * @param predicate A predicate like `array.filter(item => boolean)`
   */
  deleteMany(predicate) {
    return __async(this, null, function* () {
      const list = yield this.list();
      for (const item of list) {
        if (predicate(item)) {
          yield this.sb.delete(item.name);
          yield this.sb.delete(`${PREFIX_METADATA}${item.name}`);
        }
      }
    });
  }
  /**
   * Write the metadata of the file to disk.
   */
  writeMetadata(name, metadata) {
    return __async(this, null, function* () {
      const blob = new Blob([JSON.stringify(metadata)], { type: "text/plain" });
      yield this.sb.write(`${PREFIX_METADATA}${name}`, blob.stream());
    });
  }
};
var cache_manager_default = CacheManager;
function urlToFileName(url, prefix) {
  return __async(this, null, function* () {
    const hashBuffer = yield crypto.subtle.digest(
      "SHA-1",
      new TextEncoder().encode(url)
    );
    const hashArray = Array.from(new Uint8Array(hashBuffer));
    const hashHex = hashArray.map((b) => b.toString(16).padStart(2, "0")).join("");
    return `${prefix}${hashHex}_${url.split("/").pop()}`;
  });
}

// src/model-manager.ts
var DEFAULT_PARALLEL_DOWNLOADS = 3;
var ModelValidationStatus = /* @__PURE__ */ ((ModelValidationStatus2) => {
  ModelValidationStatus2["VALID"] = "valid";
  ModelValidationStatus2["INVALID"] = "invalid";
  ModelValidationStatus2["DELETED"] = "deleted";
  return ModelValidationStatus2;
})(ModelValidationStatus || {});
var Model = class {
  constructor(modelManager, url, mmprojUrl, savedFiles) {
    __publicField(this, "modelManager");
    /**
     * URL to the GGUF file (in case it contains multiple shards, the URL should point to the first shard)
     *
     * This URL will be used to identify the model in the cache. There can't be 2 models with the same URL.
     */
    __publicField(this, "url");
    /**
     * URL to mmproj file, if exists
     */
    __publicField(this, "mmprojUrl");
    /**
     * Size in bytes (total size of all shards).
     *
     * A value of -1 means the model is deleted from the cache. You must call `ModelManager.downloadModel` to re-download the model.
     */
    __publicField(this, "size");
    /**
     * List of all shards in the cache, sorted by original URL (ascending order)
     */
    __publicField(this, "files");
    this.modelManager = modelManager;
    this.url = url;
    this.mmprojUrl = mmprojUrl;
    if (savedFiles) {
      this.files = this.getAllFiles(savedFiles);
      this.size = sumArr(this.files.map((f) => f.metadata.originalSize));
    } else {
      this.files = [];
      this.size = 0;
    }
  }
  /**
   * Open and get a list of all shards as Blobs
   */
  open() {
    return __async(this, null, function* () {
      if (this.size === -1) {
        throw new WllamaError(
          `Model is deleted from the cache; Call ModelManager.downloadModel to re-download the model`,
          "load_error"
        );
      }
      const blobs = [];
      for (const file of this.files) {
        const blob = yield this.modelManager.cacheManager.open(file.name);
        if (!blob) {
          throw new Error(
            `Failed to open file ${file.name}; Hint: the model may be invalid, please refresh it`
          );
        }
        blobs.push(blob);
      }
      return blobs;
    });
  }
  /**
   * Validate the model files.
   *
   * If the model is invalid, the model manager will not be able to use it. You must call `refresh` to re-download the model.
   *
   * Cases that model is invalid:
   * - The model is deleted from the cache
   * - The model files are missing (or the download is interrupted)
   */
  validate() {
    let nbShards = ModelManager.parseModelUrl(this.url).length;
    if (this.mmprojUrl) {
      nbShards += 1;
    }
    if (this.size === -1) {
      return "deleted" /* DELETED */;
    }
    if (this.size < 16 || this.files.length !== nbShards) {
      return "invalid" /* INVALID */;
    }
    for (const file of this.files) {
      if (!file.metadata || file.metadata.originalSize !== file.size) {
        return "invalid" /* INVALID */;
      }
    }
    return "valid" /* VALID */;
  }
  /**
   * In case the model is invalid, call this function to re-download the model
   */
  refresh() {
    return __async(this, arguments, function* (options = {}) {
      var _a;
      const urls = ModelManager.parseModelUrl(this.url);
      if (this.mmprojUrl) {
        urls.push(this.mmprojUrl);
      }
      const works = urls.map((url, index) => ({
        url,
        index
      }));
      this.modelManager.logger.debug("Downloading model files:", urls);
      const nParallel = (_a = this.modelManager.params.parallelDownloads) != null ? _a : DEFAULT_PARALLEL_DOWNLOADS;
      const totalSize = yield this.getTotalDownloadSize(urls);
      const loadedSize = [];
      const worker = () => __async(this, null, function* () {
        while (works.length > 0) {
          const w = works.shift();
          if (!w) break;
          yield this.modelManager.cacheManager.download(w.url, __spreadProps(__spreadValues({}, options), {
            metadataAdditional: {
              originalURL: w.url,
              mmprojURL: this.mmprojUrl
            },
            progressCallback: ({ loaded }) => {
              var _a2;
              loadedSize[w.index] = loaded;
              (_a2 = options.progressCallback) == null ? void 0 : _a2.call(options, {
                loaded: sumArr(loadedSize),
                total: totalSize
              });
            }
          }));
        }
      });
      const promises = [];
      for (let i = 0; i < nParallel; i++) {
        promises.push(worker());
        loadedSize.push(0);
      }
      yield Promise.all(promises);
      this.files = this.getAllFiles(yield this.modelManager.cacheManager.list());
      this.size = this.files.reduce((acc, f) => acc + f.metadata.originalSize, 0);
    });
  }
  /**
   * Remove the model from the cache
   */
  remove() {
    return __async(this, null, function* () {
      this.files = this.getAllFiles(yield this.modelManager.cacheManager.list());
      yield this.modelManager.cacheManager.deleteMany(
        (f) => !!this.files.find((file) => file.name === f.name)
      );
      this.size = -1;
    });
  }
  getAllFiles(savedFiles) {
    const allUrls = new Set(ModelManager.parseModelUrl(this.url));
    if (this.mmprojUrl) {
      allUrls.add(this.mmprojUrl);
    }
    const allFiles = [];
    for (const url of allUrls) {
      const file = savedFiles.find((f) => f.metadata.originalURL === url);
      if (!file) {
        throw new Error(`Model file not found: ${url}`);
      }
      allFiles.push(file);
    }
    allFiles.sort(
      (a, b) => a.metadata.originalURL.localeCompare(b.metadata.originalURL)
    );
    return allFiles;
  }
  getTotalDownloadSize(urls) {
    return __async(this, null, function* () {
      const responses = yield Promise.all(
        urls.map((url) => fetch(url, { method: "HEAD" }))
      );
      const sizes = responses.map(
        (res) => Number(res.headers.get("content-length") || "0")
      );
      return sumArr(sizes);
    });
  }
};
var ModelManager = class _ModelManager {
  constructor(params = {}) {
    // The CacheManager singleton, can be accessed by user
    __publicField(this, "cacheManager");
    __publicField(this, "params");
    __publicField(this, "logger");
    this.cacheManager = params.cacheManager || new cache_manager_default();
    this.params = params;
    this.logger = params.logger || console;
  }
  /**
   * Parses a model URL and returns an array of URLs based on the following patterns:
   * - If the input URL is an array, it returns the array itself.
   * - If the input URL is a string in the `gguf-split` format, it returns an array containing the URL of each shard in ascending order.
   * - Otherwise, it returns an array containing the input URL as a single element array.
   * @param modelUrl URL or list of URLs
   */
  static parseModelUrl(modelUrl) {
    var _a;
    if (Array.isArray(modelUrl)) {
      return modelUrl;
    }
    const urlPartsRegex = /-(\d{5})-of-(\d{5})\.gguf(?:\?.*)?$/;
    const queryMatch = modelUrl.match(/\.gguf(\?.*)?$/);
    const queryParams = (_a = queryMatch == null ? void 0 : queryMatch[1]) != null ? _a : "";
    const matches = modelUrl.match(urlPartsRegex);
    if (!matches) {
      return [modelUrl];
    }
    const baseURL = modelUrl.replace(urlPartsRegex, "");
    const total = matches[2];
    const paddedShardIds = Array.from(
      { length: Number(total) },
      (_, index) => (index + 1).toString().padStart(5, "0")
    );
    return paddedShardIds.map(
      (current) => `${baseURL}-${current}-of-${total}.gguf${queryParams}`
    );
  }
  /**
   * Get all models in the cache
   */
  getModels() {
    return __async(this, arguments, function* (opts = {}) {
      const cachedFiles = yield this.cacheManager.list();
      let models = [];
      for (const file of cachedFiles) {
        if (!file.metadata.originalURL) continue;
        const shards = _ModelManager.parseModelUrl(file.metadata.originalURL);
        const mmprojUrl = file.metadata.mmprojURL;
        const isFirstShard = shards.length === 1 || shards[0] === file.metadata.originalURL;
        if (isFirstShard) {
          models.push(
            new Model(this, file.metadata.originalURL, mmprojUrl, cachedFiles)
          );
        }
      }
      if (!opts.includeInvalid) {
        models = models.filter(
          (m) => m.validate() === "valid" /* VALID */
        );
      }
      return models;
    });
  }
  /**
   * Download a model from the given URL.
   *
   * The URL must end with `.gguf`
   */
  downloadModel(_0) {
    return __async(this, arguments, function* (sourceOrURL, options = {}) {
      const source = isString(sourceOrURL) ? { url: sourceOrURL } : sourceOrURL;
      if (!isValidGgufFile(source.url)) {
        throw new WllamaError(
          `Invalid model URL: ${source.url}; URL must ends with ".gguf"`,
          "download_error"
        );
      }
      const model = new Model(this, source.url, source.mmprojUrl);
      const validity = model.validate();
      if (validity !== "valid" /* VALID */) {
        yield model.refresh(options);
      }
      return model;
    });
  }
  /**
   * Get a model from the cache or download it if it's not available.
   */
  getModelOrDownload(_0) {
    return __async(this, arguments, function* (source, options = {}) {
      var _a;
      const models = yield this.getModels();
      const model = models.find((m) => m.url === source.url);
      if (model) {
        (_a = options.progressCallback) == null ? void 0 : _a.call(options, { loaded: model.size, total: model.size });
        return model;
      }
      return this.downloadModel(source, options);
    });
  }
  /**
   * Remove all models from the cache
   */
  clear() {
    return __async(this, null, function* () {
      yield this.cacheManager.clear();
    });
  }
};

// src/types/types.ts
var LogLevel = /* @__PURE__ */ ((LogLevel2) => {
  LogLevel2[LogLevel2["DEBUG"] = 1] = "DEBUG";
  LogLevel2[LogLevel2["INFO"] = 2] = "INFO";
  LogLevel2[LogLevel2["WARN"] = 3] = "WARN";
  LogLevel2[LogLevel2["ERROR"] = 4] = "ERROR";
  return LogLevel2;
})(LogLevel || {});

// src/wasm-from-cdn.ts
var WasmCompatFromCDN = {
  worker: "https://cdn.jsdelivr.net/npm/@wllama/wllama-compat@3.6.0/wasm/wllama.js",
  wasm: "https://cdn.jsdelivr.net/npm/@wllama/wllama-compat@3.6.0/wasm/wllama.wasm"
};

// src/wllama.ts
var LoggerWithoutDebug = __spreadProps(__spreadValues({}, console), {
  debug: () => {
  }
});
var WllamaError = class extends Error {
  constructor(message, type = "unknown_error") {
    super(message);
    __publicField(this, "type");
    this.type = type;
  }
};
var WllamaAbortError = class extends Error {
  constructor() {
    super("Operation aborted");
    __publicField(this, "name", "AbortError");
  }
};
var WllamaRuntimeError = class extends Error {
  constructor(message, stack) {
    super(message);
    __publicField(this, "name", "RuntimeError");
    __publicField(this, "stack");
    this.stack = stack;
  }
};
var Wllama = class {
  constructor(pathConfig, wllamaConfig = {}) {
    // The CacheManager and ModelManager are singleton, can be accessed by user
    __publicField(this, "cacheManager");
    __publicField(this, "modelManager");
    __publicField(this, "compat", null);
    __publicField(this, "proxy", null);
    __publicField(this, "config");
    __publicField(this, "pathConfig");
    __publicField(this, "useMultiThread", false);
    __publicField(this, "nbThreads", 1);
    __publicField(this, "useEmbeddings", false);
    __publicField(this, "useRerank", false);
    // available when loaded
    __publicField(this, "loadedContextInfo", null);
    __publicField(this, "seed");
    __publicField(this, "bosToken", -1);
    __publicField(this, "eosToken", -1);
    __publicField(this, "eotToken", -1);
    __publicField(this, "eogTokens", /* @__PURE__ */ new Set());
    __publicField(this, "addBosToken", false);
    __publicField(this, "addEosToken", false);
    __publicField(this, "mediaMarker");
    __publicField(this, "chatTemplate");
    __publicField(this, "metadata");
    __publicField(this, "hasEncoder", false);
    __publicField(this, "decoderStartToken", -1);
    // note: we overlay instead of using llama-server default_template_kwargs, because we cannot transfer complex data structure via GLUE
    // overlay allow mixed data type or nested structure for kwargs
    __publicField(this, "chatTemplateKwargs", {});
    var _a, _b, _c;
    checkEnvironmentCompatible();
    if (!pathConfig) throw new WllamaError("AssetsPathConfig is required");
    this.pathConfig = pathConfig;
    this.config = wllamaConfig;
    this.cacheManager = (_a = wllamaConfig.cacheManager) != null ? _a : new cache_manager_default();
    this.modelManager = (_c = wllamaConfig.modelManager) != null ? _c : new ModelManager({
      cacheManager: this.cacheManager,
      logger: (_b = wllamaConfig.logger) != null ? _b : console,
      parallelDownloads: wllamaConfig.parallelDownloads,
      allowOffline: wllamaConfig.allowOffline
    });
    this.setCompat("default");
  }
  logger() {
    var _a;
    return (_a = this.config.logger) != null ? _a : console;
  }
  checkModelLoaded() {
    if (!this.isModelLoaded()) {
      throw new WllamaError(
        "loadModel() is not yet called",
        "model_not_loaded"
      );
    }
  }
  /**
   * Get the libllama version string, e.g. "b6327-4d74393".
   *
   * @returns version string embedded at build time.
   */
  static getLibllamaVersion() {
    return LIBLLAMA_VERSION;
  }
  /**
   * Set compatibility options for Wllama.
   * @param compat Set to null to disable compatibility, or 'default' to use the default compat resources from CDN.
   * @param mode 'safari' by default; If set to 'firefox_safari', the compat mode will **also** be enabled on Firefox, which will significantly degrade the performance but allow using WebGPU on Firefox.
   */
  setCompat(compat, mode = "safari") {
    if (mode === "safari") {
      if (isFirefox()) {
        this.compat = null;
        return;
      }
    }
    this.compat = compat === "default" ? WasmCompatFromCDN : compat;
  }
  /**
   * Check if the model is loaded via `loadModel()`
   */
  isModelLoaded() {
    return !!this.proxy && !!this.metadata;
  }
  /**
   * Get token ID associated to BOS (begin of sentence) token.
   *
   * NOTE: This can only being used after `loadModel` is called.
   *
   * @returns -1 if the model is not loaded.
   */
  getBOS() {
    return this.bosToken;
  }
  /**
   * Get token ID associated to EOS (end of sentence) token.
   *
   * NOTE: This can only being used after `loadModel` is called.
   *
   * @returns -1 if the model is not loaded.
   */
  getEOS() {
    return this.eosToken;
  }
  /**
   * Get token ID associated to EOT (end of turn) token.
   *
   * NOTE: This can only being used after `loadModel` is called.
   *
   * @returns -1 if the model is not loaded.
   */
  getEOT() {
    return this.eotToken;
  }
  /**
   * Check if a given token is end-of-generation token (e.g. EOS, EOT, etc.)
   *
   * @param token the token ID to be checked
   * @returns true if the token is EOS, EOT, or any other end-of-generation tokens
   */
  isTokenEOG(token) {
    return token === this.eosToken || token === this.eotToken || this.eogTokens.has(token);
  }
  /**
   * Get token ID associated to token used by decoder, to start generating output sequence(only usable for encoder-decoder architecture). In other words, encoder uses normal BOS and decoder uses this token.
   *
   * NOTE: This can only being used after `loadModel` is called.
   *
   * @returns -1 if the model is not loaded.
   */
  getDecoderStartToken() {
    return this.decoderStartToken;
  }
  /**
   * Get model hyper-parameters and metadata
   *
   * NOTE: This can only being used after `loadModel` is called.
   *
   * @returns ModelMetadata
   */
  getModelMetadata() {
    this.checkModelLoaded();
    return this.metadata;
  }
  /**
   * Check if we're currently using multi-thread build.
   *
   * NOTE: This can only being used after `loadModel` is called.
   *
   * @returns true if multi-thread is used.
   */
  isMultithread() {
    this.checkModelLoaded();
    return this.useMultiThread;
  }
  /**
   * Get number of threads used in the current context.
   *
   * NOTE: This can only being used after `loadModel` is called.
   *
   * @returns number of threads
   */
  getNumThreads() {
    this.checkModelLoaded();
    return this.useMultiThread ? this.nbThreads : 1;
  }
  /**
   * Check if the current model uses encoder-decoder architecture
   *
   * NOTE: This can only being used after `loadModel` is called.
   *
   * @returns true if multi-thread is used.
   */
  isEncoderDecoderArchitecture() {
    this.checkModelLoaded();
    return this.hasEncoder;
  }
  /**
   * Must we add BOS token to the tokenized sequence?
   *
   * NOTE: This can only being used after `loadModel` is called.
   *
   * @returns true if BOS token must be added to the sequence
   */
  mustAddBosToken() {
    this.checkModelLoaded();
    return this.addBosToken;
  }
  /**
   * Must we add EOS token to the tokenized sequence?
   *
   * NOTE: This can only being used after `loadModel` is called.
   *
   * @returns true if EOS token must be added to the sequence
   */
  mustAddEosToken() {
    this.checkModelLoaded();
    return this.addEosToken;
  }
  /**
   * Get the jinja chat template comes with the model. It only available if the original model (before converting to gguf) has the template in `tokenizer_config.json`
   *
   * NOTE: This can only being used after `loadModel` is called.
   *
   * @returns the jinja template. null if there is no template in gguf
   */
  getChatTemplate() {
    var _a;
    this.checkModelLoaded();
    return (_a = this.chatTemplate) != null ? _a : null;
  }
  /**
   * Check if WebGPU is supported by the current environment.
   * @returns true if WebGPU is supported
   */
  isSupportWebGPU() {
    return isSupportWebGPU();
  }
  /**
   * Load model from a given URL (or a list of URLs, in case the model is splitted into smaller files)
   * - If the model already been downloaded (via `downloadModel()`), then we will use the cached model
   * - Else, we download the model from internet
   * @param modelSourceOrURL
   * @param params
   */
  loadModelFromUrl(_0) {
    return __async(this, arguments, function* (modelSourceOrURL, params = {}) {
      var _a;
      const source = isString(modelSourceOrURL) ? { url: modelSourceOrURL } : modelSourceOrURL;
      const useCache = (_a = params.useCache) != null ? _a : true;
      const model = useCache ? yield this.modelManager.getModelOrDownload(source, params) : yield this.modelManager.downloadModel(source, params);
      const blobs = yield model.open();
      return yield this.loadModel(blobs, params);
    });
  }
  /**
   * Load model from a given Hugging Face model ID and file path.
   *
   * @param hfOptions
   * @param params
   */
  loadModelFromHF(_0) {
    return __async(this, arguments, function* (hfOptions, params = {}) {
      const source = yield getHFModelSource(hfOptions);
      return yield this.loadModelFromUrl(source, params);
    });
  }
  /**
   * Load model from a given list of Blob.
   *
   * You can pass multiple buffers into the function (in case the model contains multiple shards).
   *
   * @param ggufBlobsOrModel Can be either list of Blobs (in case you use local file), or a Model object (in case you use ModelManager)
   * @param params LoadModelParams
   */
  loadModel(_0) {
    return __async(this, arguments, function* (ggufBlobsOrModel, params = {}) {
      var _a, _b, _c, _d, _e, _f, _g, _h, _i, _j;
      const blobs = ggufBlobsOrModel instanceof Model ? yield ggufBlobsOrModel.open() : [...ggufBlobsOrModel];
      if (blobs.some((b) => b.size === 0)) {
        throw new WllamaError(
          "Input model (or splits) must be non-empty Blob or File",
          "load_error"
        );
      }
      if (!this.pathConfig["default"]) {
        throw new WllamaError(
          '"default" is missing from pathConfig',
          "load_error"
        );
      }
      if (this.proxy) {
        throw new WllamaError("Module is already initialized", "load_error");
      }
      const supportMultiThread = yield isSupportMultiThread();
      const hwConccurency = Math.floor((navigator.hardwareConcurrency || 1) / 2);
      const nbThreads = (_a = params.n_threads) != null ? _a : hwConccurency;
      this.nbThreads = nbThreads;
      this.useMultiThread = supportMultiThread && nbThreads > 1;
      const workerResources = this.getWorkerResources();
      if (params.n_gpu_layers === 0) {
        workerResources.noWebGPU = true;
      }
      this.proxy = new ProxyToWorker(
        workerResources,
        this.useMultiThread ? nbThreads : 0,
        // 0 means disable pthread
        (_b = this.config.suppressNativeLog) != null ? _b : false,
        this.logger()
      );
      let logLevel = (_c = params.log_level) != null ? _c : 2 /* INFO */;
      if (this.config.suppressNativeLog) {
        logLevel = 9999;
      }
      const modelFiles = yield prepareBlobs(blobs);
      yield this.proxy.moduleInit(modelFiles.all);
      this.logger().debug("Calling wllamaStart...");
      const startResult = yield this.proxy.wllamaStart();
      if (!startResult.success) {
        throw new WllamaError(
          `Error while calling start function, result = ${startResult}`
        );
      }
      this.logger().debug("Loading model...");
      const loadResult = yield this.proxy.wllamaAction("load", {
        _name: "load_req",
        log_level: logLevel,
        // if async read is not supported, use mmap; refer to README-dev.md for more details
        use_mmap: !canUseAsyncFileRead(workerResources.compat),
        use_mlock: false,
        n_gpu_layers: (_d = params.n_gpu_layers) != null ? _d : 99999,
        n_ctx: (_e = params.n_ctx) != null ? _e : 1024,
        n_threads: this.useMultiThread ? nbThreads : 1,
        n_ctx_auto: false,
        // not supported for now
        mmproj_path: modelFiles.mmproj ? `/models/${MMPROJ_FILE_NAME}` : void 0,
        model_paths: modelFiles.llm.map((f) => `models/${f.name}`),
        embeddings: params.embeddings,
        offload_kqv: params.offload_kqv,
        n_batch: params.n_batch,
        n_ubatch: params.n_ubatch,
        pooling_type: params.pooling_type,
        rope_scaling_type: params.rope_scaling_type,
        rope_freq_base: params.rope_freq_base,
        rope_freq_scale: params.rope_freq_scale,
        yarn_ext_factor: params.yarn_ext_factor,
        yarn_attn_factor: params.yarn_attn_factor,
        yarn_beta_fast: params.yarn_beta_fast,
        yarn_beta_slow: params.yarn_beta_slow,
        yarn_orig_ctx: params.yarn_orig_ctx,
        cache_type_k: params.cache_type_k,
        cache_type_v: params.cache_type_v,
        // with unified KV, all sequences share one n_ctx cache, so each request can still use the full context
        n_parallel: (_f = params.n_parallel) != null ? _f : 4,
        kv_unified: (_g = params.kv_unified) != null ? _g : true,
        flash_attn: params.flash_attn,
        swa_full: params.swa_full,
        chat_template: params.chat_template,
        jinja: params.jinja,
        reasoning: params.reasoning,
        image_min_tokens: params.image_min_tokens,
        image_max_tokens: params.image_max_tokens,
        warmup: params.warmup,
        no_kv_offload: params.no_kv_offload,
        mmproj_offload: params.mmproj_offload,
        cont_batching: params.cont_batching,
        n_keep: params.n_keep,
        ctx_shift: params.ctx_shift,
        cache_idle_slots: params.cache_idle_slots,
        n_cache_reuse: params.n_cache_reuse,
        lora_paths: (_h = params.lora_adapters) == null ? void 0 : _h.map((a) => a.path),
        lora_scales: (_i = params.lora_adapters) == null ? void 0 : _i.map((a) => {
          var _a2;
          return (_a2 = a.scale) != null ? _a2 : 1;
        }),
        lora_init_without_apply: params.lora_init_without_apply,
        spec_draft_model: params.spec_draft_model,
        spec_draft_ngl: params.spec_draft_ngl,
        spec_draft_n_max: params.spec_draft_n_max,
        spec_draft_n_min: params.spec_draft_n_min,
        spec_draft_p_min: params.spec_draft_p_min,
        spec_draft_threads: params.spec_draft_threads,
        spec_draft_threads_batch: params.spec_draft_threads_batch,
        kv_overrides_keys: params.kv_overrides ? Object.keys(params.kv_overrides) : void 0,
        kv_overrides_vals: params.kv_overrides ? Object.values(params.kv_overrides) : void 0,
        reasoning_budget_tokens: params.reasoning_budget_tokens,
        reasoning_budget_message: params.reasoning_budget_message,
        reasoning_format: params.reasoning_format,
        skip_chat_parsing: params.skip_chat_parsing,
        prefill_assistant: params.prefill_assistant
      });
      const loadedCtxInfo = __spreadProps(__spreadValues({}, loadResult), {
        metadata: {}
      });
      for (let i = 0; i < loadResult.metadata_key.length; i++) {
        loadedCtxInfo.metadata[loadResult.metadata_key[i]] = loadResult.metadata_val[i];
      }
      this.seed = params.seed;
      this.bosToken = loadedCtxInfo.token_bos;
      this.eosToken = loadedCtxInfo.token_eos;
      this.eotToken = loadedCtxInfo.token_eot;
      this.useEmbeddings = !!params.embeddings;
      this.useRerank = params.pooling_type == "rank";
      this.metadata = {
        hparams: {
          nVocab: loadedCtxInfo.n_vocab,
          nCtxTrain: loadedCtxInfo.n_ctx_train,
          nEmbd: loadedCtxInfo.n_embd,
          nLayer: loadedCtxInfo.n_layer
        },
        meta: loadedCtxInfo.metadata
      };
      this.hasEncoder = !!loadedCtxInfo.has_encoder;
      this.decoderStartToken = loadedCtxInfo.token_decoder_start;
      this.addBosToken = loadedCtxInfo.add_bos_token;
      this.addEosToken = loadedCtxInfo.add_eos_token;
      this.chatTemplate = loadedCtxInfo.metadata["tokenizer.chat_template"];
      this.loadedContextInfo = loadedCtxInfo;
      this.eogTokens = new Set(loadedCtxInfo.list_tokens_eog);
      this.mediaMarker = loadedCtxInfo.media_marker;
      this.chatTemplateKwargs = (_j = params.default_template_kwargs) != null ? _j : {};
      this.logger().debug({ loadedCtxInfo });
    });
  }
  getLoadedContextInfo() {
    this.checkModelLoaded();
    if (!this.loadedContextInfo) {
      throw new WllamaError("Loaded context info is not available");
    }
    return __spreadValues({}, this.loadedContextInfo);
  }
  //////////////////////////////////////////////
  // High level API
  /**
   * Calculate embedding vector for a given text.
   * By default, BOS and EOS tokens will be added automatically. You can use the "skipBOS" and "skipEOS" option to disable it.
   * @param options OAI-compatible embedding creation options
   * @returns OAI-compatible embedding response
   */
  createEmbedding(options) {
    return __async(this, null, function* () {
      this.checkModelLoaded();
      if (!this.useEmbeddings) {
        throw new WllamaError(
          "Embeddings is not enabled. Please set it via LoadModelParams.embeddings"
        );
      }
      const result = yield this.proxy.wllamaAction(
        "embedding",
        {
          _name: "embd_req",
          data_json: JSON.stringify(options),
          files: []
          // TODO: support file input
        }
      );
      if (!result.success) {
        throw new WllamaError(
          "Model failed to start inference",
          "inference_error"
        );
      }
      return yield this.getResponse(options, false, result.req_id);
    });
  }
  /**
   * Rerank a list of documents against a query.
   * Requires the model to be loaded with embeddings: true and pooling_type: 'rank'.
   * @param options Reranking options (query, documents, top_n)
   * @returns Reranking response with relevance scores sorted highest first
   */
  createRerank(options) {
    return __async(this, null, function* () {
      var _a, _b;
      this.checkModelLoaded();
      if (!this.useEmbeddings || !this.useRerank) {
        throw new WllamaError(
          "Rerank is not enabled. Please set it via LoadModelParams: embeddings = true and pooling_type = rank"
        );
      }
      const top_n = (_a = options.top_n) != null ? _a : options.documents.length;
      let totalTokens = 0;
      const rawResults = [];
      for (let i = 0; i < options.documents.length; i++) {
        const result = yield this.proxy.wllamaAction("rerank", {
          _name: "rrnk_req",
          data_json: JSON.stringify({
            query: options.query,
            document: options.documents[i]
          })
        });
        if (!result.success) {
          throw new WllamaError(
            "Model failed to start reranking",
            "inference_error"
          );
        }
        const { score, tokens_evaluated } = yield this.getRerankResult(
          result.req_id
        );
        totalTokens += tokens_evaluated;
        rawResults.push({ index: i, score });
      }
      rawResults.sort((a, b) => b.score - a.score);
      return {
        model: (_b = this.getModelMetadata().meta["general.name"]) != null ? _b : "",
        object: "list",
        usage: { prompt_tokens: totalTokens, total_tokens: totalTokens },
        results: rawResults.slice(0, top_n).map(({ index, score }) => ({
          index,
          relevance_score: score
        }))
      };
    });
  }
  createChatCompletion(options) {
    return __async(this, null, function* () {
      var _a;
      if (Object.keys(this.chatTemplateKwargs).length > 0) {
        options = __spreadProps(__spreadValues({}, options), {
          chat_template_kwargs: __spreadValues(__spreadValues({}, this.chatTemplateKwargs), (_a = options.chat_template_kwargs) != null ? _a : {})
        });
      }
      if (options.stream && options.onData) {
        yield this.createCompletionImpl(options);
      } else if (options.stream) {
        return yield this.createCompletionGenerator(options);
      } else {
        return yield this.createCompletionImpl(__spreadProps(__spreadValues({}, options), { stream: false }));
      }
    });
  }
  createCompletion(options) {
    return __async(this, null, function* () {
      if (options.stream && options.onData) {
        yield this.createCompletionImpl(options);
      } else if (options.stream) {
        return yield this.createCompletionGenerator(options);
      } else {
        return yield this.createCompletionImpl(__spreadProps(__spreadValues({}, options), { stream: false }));
      }
    });
  }
  /**
   * Private implementation of createCompletion
   */
  createCompletionImpl(options) {
    return __async(this, null, function* () {
      this.checkModelLoaded();
      const isStream = !!options.stream;
      const isChat = !!options.messages;
      const customOpt = {};
      if (this.seed !== void 0) {
        customOpt.seed = this.seed;
      }
      let files = [];
      if (isChat) {
        const tmp = this.prepareMultimodalInput(
          options
        );
        options = tmp.params;
        files = tmp.files;
      }
      const result = yield this.proxy.wllamaAction(
        "completion",
        {
          _name: "cmpl_req",
          is_chat: isChat,
          data_json: JSON.stringify(__spreadValues(__spreadValues({}, options), customOpt)),
          files: files.map((f) => new Uint8Array(f))
        }
      );
      if (!result.success) {
        throw new WllamaError(
          "Model failed to start inference",
          "inference_error"
        );
      }
      return yield this.getResponse(
        options,
        isStream,
        result.req_id
      );
    });
  }
  /**
   * Same with `createCompletion`, but returns an async iterator instead.
   * Only called when stream=true and no onData is provided.
   */
  createCompletionGenerator(options) {
    return new Promise((resolve) => {
      const createGenerator = cbToAsyncIter(
        (callback) => {
          this.createCompletionImpl(__spreadProps(__spreadValues({}, options), {
            onData: (chunk) => callback(chunk)
          })).then(() => callback(void 0, true)).catch((err) => callback(void 0, false, err));
        }
      );
      resolve(createGenerator());
    });
  }
  /**
   * Whether the currently loaded model supports a specific input modality (e.g. image or audio).
   * @param modality
   * @returns
   */
  supportInputModality(modality) {
    this.checkModelLoaded();
    if (modality === "image") {
      return !!this.loadedContextInfo.has_image_input;
    } else if (modality === "audio") {
      return !!this.loadedContextInfo.has_audio_input;
    } else {
      throw new WllamaError(
        "Unsupported modality: " + modality,
        "unknown_error"
      );
    }
  }
  /**
   * Unload the model and free all memory.
   *
   * Note: This function will NOT crash if model is not yet loaded
   */
  exit() {
    return __async(this, null, function* () {
      var _a;
      yield (_a = this.proxy) == null ? void 0 : _a.wllamaExit();
      this.proxy = null;
    });
  }
  /**
   * [FOR DEBUGGING ONLY] Run ggml backend ops tests without loading any model.
   *
   * Initializes the wasm runtime, executes `test-backend-ops` with the given args, then shuts down.
   *
   * For more info, please refer to guides/debug.md
   *
   * @param args Arguments forwarded to test-backend-ops (e.g. ["-o", "ADD"])
   * @returns retcode (0 = all tests passed) and success flag
   */
  testBackendOps() {
    return __async(this, arguments, function* (args = []) {
      var _a;
      if (!this.pathConfig["default"]) {
        throw new WllamaError(
          '"default" is missing from pathConfig',
          "load_error"
        );
      }
      if (!(yield isSupportMultiThread())) {
        throw new WllamaError(
          "Multi-threading is required to run backend ops tests, but it is not supported in the current environment."
        );
      }
      const tmpProxy = new ProxyToWorker(
        this.getWorkerResources(),
        0,
        // single-thread; no model needed
        (_a = this.config.suppressNativeLog) != null ? _a : false,
        this.logger()
      );
      try {
        yield tmpProxy.moduleInit([]);
        const startResult = yield tmpProxy.wllamaStart();
        if (!startResult.success) {
          throw new WllamaError(
            `Error while calling start function, result = ${startResult}`
          );
        }
        const result = yield tmpProxy.wllamaAction(
          "test_backend_ops",
          { _name: "tbop_req", args: ["test-backend-ops", ...args] }
        );
        return { retcode: result.retcode, success: result.success };
      } finally {
        yield tmpProxy.wllamaExit();
      }
    });
  }
  //////////////////////////////////////////////
  // Low level API
  // TODO: add back
  /**
   * get debug info
   */
  _getDebugInfo() {
    return __async(this, null, function* () {
      this.checkModelLoaded();
      return yield this.proxy.wllamaDebug();
    });
  }
  //////////////////////////////////////////////
  // Utils
  jsonDecode(data_json) {
    try {
      return JSON.parse(data_json);
    } catch (e) {
      this.logger().error("Failed to parse JSON:", data_json);
      throw new WllamaError("Failed to parse model output", "inference_error");
    }
  }
  prepareMultimodalInput(params) {
    const msg = params.messages;
    const msgNew = [];
    const files = [];
    for (const m of msg) {
      if (Array.isArray(m.content)) {
        const newContent = [];
        for (const c of m.content) {
          if (c.type === "text") {
            newContent.push(c);
          } else {
            if (!this.mediaMarker) {
              throw new WllamaError(
                "Media marker is undefined",
                "inference_error"
              );
            }
            files.push(c.data);
            newContent.push({
              type: "text",
              text: this.mediaMarker
            });
          }
        }
        msgNew.push(__spreadProps(__spreadValues({}, m), {
          content: newContent
        }));
      } else {
        msgNew.push(m);
      }
    }
    return {
      params: __spreadProps(__spreadValues({}, params), {
        messages: msgNew
      }),
      files
    };
  }
  // release the slot occupied by the request; cancelling an already-finished request is a no-op
  cancelRequest(reqId) {
    return __async(this, null, function* () {
      try {
        yield this.proxy.wllamaAction("cancel", {
          _name: "cncl_req",
          req_id: reqId
        });
      } catch (e) {
        this.logger().warn("Failed to cancel request", reqId, e);
      }
    });
  }
  getRerankResult(reqId) {
    return __async(this, null, function* () {
      let completed = false;
      try {
        while (true) {
          const chunk = yield this.proxy.wllamaAction(
            "get_result",
            { _name: "gres_req", req_id: reqId }
          );
          const jsonString = chunk.data_json;
          if (jsonString && jsonString.length > 0) {
            if (chunk.is_error) {
              const jsonData = this.jsonDecode(jsonString);
              throw new WllamaError(
                jsonData.message || "Unknown reranking error",
                "inference_error"
              );
            }
            completed = true;
            return this.jsonDecode(jsonString);
          }
          if (!chunk.has_more) {
            completed = true;
            break;
          }
        }
        throw new WllamaError("No reranking result received", "inference_error");
      } finally {
        if (!completed) {
          yield this.cancelRequest(reqId);
        }
      }
    });
  }
  getResponse(options, isStream, reqId) {
    return __async(this, null, function* () {
      var _a, _b;
      let finalResult = null;
      let completed = false;
      try {
        while (true) {
          if ((_a = options.abortSignal) == null ? void 0 : _a.aborted) {
            throw new WllamaAbortError();
          }
          const result_chunk = yield this.proxy.wllamaAction(
            "get_result",
            {
              _name: "gres_req",
              req_id: reqId
            }
          );
          const jsonString = result_chunk.data_json;
          if (!jsonString || jsonString.length === 0) {
            if (!result_chunk.has_more) {
              completed = true;
              break;
            } else {
              continue;
            }
          }
          if (jsonString == "null") {
            continue;
          }
          let jsonData = this.jsonDecode(jsonString);
          finalResult = jsonData;
          if (result_chunk.is_error) {
            this.logger().error("Model returned an error:", jsonData);
            throw new WllamaError(
              jsonData.message || "Unknown inference error",
              "inference_error"
            );
          }
          if (isStream) {
            if (!Array.isArray(jsonData)) {
              jsonData = [jsonData];
            }
            for (const chunk of jsonData) {
              (_b = options.onData) == null ? void 0 : _b.call(options, chunk);
              finalResult = chunk;
            }
          }
          if (!result_chunk.has_more) {
            completed = true;
            break;
          }
        }
      } finally {
        if (!completed) {
          yield this.cancelRequest(reqId);
        }
      }
      return finalResult;
    });
  }
  getWorkerResources() {
    const workerResources = {
      wasmPath: absoluteUrl(this.pathConfig["default"]),
      compat: false
    };
    if (needCompat()) {
      if (!this.compat) {
        this.logger().warn(
          "Not using compat mode" + (isFirefox() ? " (expected on Firefox - WebGPU will be disabled)" : "")
        );
      } else {
        const isUsingDefault = this.compat.worker === WasmCompatFromCDN.worker && this.compat.wasm === WasmCompatFromCDN.wasm;
        if (isUsingDefault) {
          this.logger().warn(
            "Compatibility mode is activated, using resources from CDN. To use local resources, please refer to @wllama/wllama-compat package."
          );
          this.logger().warn(
            "IMPORTANT: Performance will be significantly degraded in compatibility mode."
          );
        }
        workerResources.wasmPath = absoluteUrl(this.compat.wasm);
        workerResources.jsPath = this.compat.worker;
        workerResources.compat = true;
      }
    }
    if (isFirefox()) {
      if (workerResources.compat) {
        this.logger().warn(
          'Using compat mode on Firefox, performance will be significantly degraded; Consider enabling "javascript.options.wasm_js_promise_integration" in "about:config".'
        );
      } else if (!isSupportJSPI()) {
        this.logger().warn(
          'WebGPU is disabled on Firefox due to missing JSPI support. Please consider enabling compat mode, or enabling "javascript.options.wasm_js_promise_integration" in "about:config".'
        );
      }
    }
    return workerResources;
  }
};
export {
  CacheManager,
  LogLevel,
  LoggerWithoutDebug,
  Model,
  ModelManager,
  ModelValidationStatus,
  POLYFILL_ETAG,
  Wllama,
  WllamaAbortError,
  WllamaError,
  WllamaRuntimeError,
  getHFFileSHA256,
  getHFModelSource,
  isValidGgufFile
};
