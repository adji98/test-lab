/******/ (function(modules) { // webpackBootstrap
/******/ 	// The module cache
/******/ 	var installedModules = {};
/******/
/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {
/******/
/******/ 		// Check if module is in cache
/******/ 		if(installedModules[moduleId]) {
/******/ 			return installedModules[moduleId].exports;
/******/ 		}
/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = installedModules[moduleId] = {
/******/ 			i: moduleId,
/******/ 			l: false,
/******/ 			exports: {}
/******/ 		};
/******/
/******/ 		// Execute the module function
/******/ 		modules[moduleId].call(module.exports, module, module.exports, __webpack_require__);
/******/
/******/ 		// Flag the module as loaded
/******/ 		module.l = true;
/******/
/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}
/******/
/******/
/******/ 	// expose the modules object (__webpack_modules__)
/******/ 	__webpack_require__.m = modules;
/******/
/******/ 	// expose the module cache
/******/ 	__webpack_require__.c = installedModules;
/******/
/******/ 	// define getter function for harmony exports
/******/ 	__webpack_require__.d = function(exports, name, getter) {
/******/ 		if(!__webpack_require__.o(exports, name)) {
/******/ 			Object.defineProperty(exports, name, {
/******/ 				configurable: false,
/******/ 				enumerable: true,
/******/ 				get: getter
/******/ 			});
/******/ 		}
/******/ 	};
/******/
/******/ 	// define __esModule on exports
/******/ 	__webpack_require__.r = function(exports) {
/******/ 		Object.defineProperty(exports, '__esModule', { value: true });
/******/ 	};
/******/
/******/ 	// getDefaultExport function for compatibility with non-harmony modules
/******/ 	__webpack_require__.n = function(module) {
/******/ 		var getter = module && module.__esModule ?
/******/ 			function getDefault() { return module['default']; } :
/******/ 			function getModuleExports() { return module; };
/******/ 		__webpack_require__.d(getter, 'a', getter);
/******/ 		return getter;
/******/ 	};
/******/
/******/ 	// Object.prototype.hasOwnProperty.call
/******/ 	__webpack_require__.o = function(object, property) { return Object.prototype.hasOwnProperty.call(object, property); };
/******/
/******/ 	// __webpack_public_path__
/******/ 	__webpack_require__.p = "";
/******/
/******/
/******/ 	// webpack-livereload-plugin
/******/ 	(function() {
/******/ 	  if (typeof window === "undefined") { return };
/******/ 	  var id = "webpack-livereload-plugin-script";
/******/ 	  if (document.getElementById(id)) { return; }
/******/ 	  var el = document.createElement("script");
/******/ 	  el.id = id;
/******/ 	  el.async = true;
/******/ 	  el.src = "//" + location.hostname + ":35729/livereload.js";
/******/ 	  document.getElementsByTagName("head")[0].appendChild(el);
/******/ 	}());
/******/ 	// Load entry module and return exports
/******/ 	return __webpack_require__(__webpack_require__.s = 4);
/******/ })
/************************************************************************/
/******/ ({

/***/ "./assets/js/modernizr.custom.min.js":
/*!*******************************************!*\
  !*** ./assets/js/modernizr.custom.min.js ***!
  \*******************************************/
/*! no static exports found */
/***/ (function(module, exports) {

eval("function _typeof(obj) { if (typeof Symbol === \"function\" && typeof Symbol.iterator === \"symbol\") { _typeof = function _typeof(obj) { return typeof obj; }; } else { _typeof = function _typeof(obj) { return obj && typeof Symbol === \"function\" && obj.constructor === Symbol && obj !== Symbol.prototype ? \"symbol\" : typeof obj; }; } return _typeof(obj); }\n\n/* Modernizr 2.0.6 (Custom Build) | MIT & BSD\n * Build: http://www.modernizr.com/download/#-touch-teststyles-prefixes\n */\n;\n\nwindow.Modernizr = function (a, b, c) {\n  function y(a, b) {\n    return !!~(\"\" + a).indexOf(b);\n  }\n\n  function x(a, b) {\n    return _typeof(a) === b;\n  }\n\n  function w(a, b) {\n    return v(m.join(a + \";\") + (b || \"\"));\n  }\n\n  function v(a) {\n    j.cssText = a;\n  }\n\n  var d = \"2.0.6\",\n      e = {},\n      f = b.documentElement,\n      g = b.head || b.getElementsByTagName(\"head\")[0],\n      h = \"modernizr\",\n      i = b.createElement(h),\n      j = i.style,\n      k,\n      l = Object.prototype.toString,\n      m = \" -webkit- -moz- -o- -ms- -khtml- \".split(\" \"),\n      n = {},\n      o = {},\n      p = {},\n      q = [],\n      r = function r(a, c, d, e) {\n    var g,\n        i,\n        j,\n        k = b.createElement(\"div\");\n    if (parseInt(d, 10)) while (d--) {\n      j = b.createElement(\"div\"), j.id = e ? e[d] : h + (d + 1), k.appendChild(j);\n    }\n    g = [\"&shy;\", \"<style>\", a, \"</style>\"].join(\"\"), k.id = h, k.innerHTML += g, f.appendChild(k), i = c(k, a), k.parentNode.removeChild(k);\n    return !!i;\n  },\n      s,\n      t = {}.hasOwnProperty,\n      u;\n\n  !x(t, c) && !x(t.call, c) ? u = function u(a, b) {\n    return t.call(a, b);\n  } : u = function u(a, b) {\n    return b in a && x(a.constructor.prototype[b], c);\n  };\n\n  var z = function (c, d) {\n    var f = c.join(\"\"),\n        g = d.length;\n    r(f, function (c, d) {\n      var f = b.styleSheets[b.styleSheets.length - 1],\n          h = f.cssRules && f.cssRules[0] ? f.cssRules[0].cssText : f.cssText || \"\",\n          i = c.childNodes,\n          j = {};\n\n      while (g--) {\n        j[i[g].id] = i[g];\n      }\n\n      e.touch = \"ontouchstart\" in a || j.touch.offsetTop === 9;\n    }, g, d);\n  }([, [\"@media (\", m.join(\"touch-enabled),(\"), h, \")\", \"{#touch{top:9px;position:absolute}}\"].join(\"\")], [, \"touch\"]);\n\n  n.touch = function () {\n    return e.touch;\n  };\n\n  for (var A in n) {\n    u(n, A) && (s = A.toLowerCase(), e[s] = n[A](), q.push((e[s] ? \"\" : \"no-\") + s));\n  }\n\n  v(\"\"), i = k = null, e._version = d, e._prefixes = m, e.testStyles = r;\n  return e;\n}(this, this.document);\n\n//# sourceURL=webpack:///./assets/js/modernizr.custom.min.js?");

/***/ }),

/***/ 4:
/*!*************************************************!*\
  !*** multi ./assets/js/modernizr.custom.min.js ***!
  \*************************************************/
/*! no static exports found */
/***/ (function(module, exports, __webpack_require__) {

eval("module.exports = __webpack_require__(/*! ./assets/js/modernizr.custom.min.js */\"./assets/js/modernizr.custom.min.js\");\n\n\n//# sourceURL=webpack:///multi_./assets/js/modernizr.custom.min.js?");

/***/ })

/******/ });