(function(t){function e(e){for(var n,l,i=e[0],u=e[1],s=e[2],c=0,f=[];c<i.length;c++)l=i[c],Object.prototype.hasOwnProperty.call(a,l)&&a[l]&&f.push(a[l][0]),a[l]=0;for(n in u)Object.prototype.hasOwnProperty.call(u,n)&&(t[n]=u[n]);p&&p(e);while(f.length)f.shift()();return o.push.apply(o,s||[]),r()}function r(){for(var t,e=0;e<o.length;e++){for(var r=o[e],n=!0,i=1;i<r.length;i++){var u=r[i];0!==a[u]&&(n=!1)}n&&(o.splice(e--,1),t=l(l.s=r[0]))}return t}var n={},a={app:0},o=[];function l(e){if(n[e])return n[e].exports;var r=n[e]={i:e,l:!1,exports:{}};return t[e].call(r.exports,r,r.exports,l),r.l=!0,r.exports}l.m=t,l.c=n,l.d=function(t,e,r){l.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:r})},l.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},l.t=function(t,e){if(1&e&&(t=l(t)),8&e)return t;if(4&e&&"object"===typeof t&&t&&t.__esModule)return t;var r=Object.create(null);if(l.r(r),Object.defineProperty(r,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var n in t)l.d(r,n,function(e){return t[e]}.bind(null,n));return r},l.n=function(t){var e=t&&t.__esModule?function(){return t["default"]}:function(){return t};return l.d(e,"a",e),e},l.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},l.p="/assets/app/";var i=window["webpackJsonp"]=window["webpackJsonp"]||[],u=i.push.bind(i);i.push=e,i=i.slice();for(var s=0;s<i.length;s++)e(i[s]);var p=u;o.push([0,"chunk-vendors"]),r()})({0:function(t,e,r){t.exports=r("56d7")},"034f":function(t,e,r){"use strict";r("85ec")},"56d7":function(t,e,r){"use strict";r.r(e);var n=r("2b0e"),a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{attrs:{id:"app"}},[n("img",{attrs:{alt:"Vue logo",src:r("cf05")}}),n("HelloWorld",{attrs:{msg:"Welcome to Lambda Platform"}})],1)},o=[],l=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",{staticClass:"hello"},[r("h1",[t._v(t._s(t.msg))]),t._m(0),r("h3",[t._v("What is included")]),t._m(1)])},i=[function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("p",[t._v(" Go to system control, please "),r("a",{attrs:{href:"/auth/login",target:"_blank"}},[t._v("sign in")]),t._v(". ")])},function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("ul",[r("li",[r("a",{attrs:{href:"https://github.com/lambda-platform/lambda",target:"_blank",rel:"noopener"}},[t._v("Lambda")])]),r("li",[r("a",{attrs:{href:"https://github.com/lambda-platform/puzzle",target:"_blank",rel:"noopener"}},[t._v("Puzzle")])]),r("li",[r("a",{attrs:{href:"https://github.com/lambda-platform/agent",target:"_blank",rel:"noopener"}},[t._v("Agent")])]),r("li",[r("a",{attrs:{href:"https://github.com/lambda-platform/datasource",target:"_blank",rel:"noopener"}},[t._v("Data Source")])]),r("li",[r("a",{attrs:{href:"https://github.com/lambda-platform/krud",target:"_blank",rel:"noopener"}},[t._v("Krud")])]),r("li",[r("a",{attrs:{href:"https://github.com/lambda-platform/dataform",target:"_blank",rel:"noopener"}},[t._v("Data From")])]),r("li",[r("a",{attrs:{href:"https://github.com/lambda-platform/datagrid",target:"_blank",rel:"noopener"}},[t._v("Data Grid")])]),r("li",[t._v("App example")]),r("li",[t._v("App Admin example (optional)")])])}],u={name:"HelloWorld",props:{msg:String}},s=u,p=(r("b132"),r("2877")),c=Object(p["a"])(s,l,i,!1,null,"59039e0c",null),f=c.exports,b={name:"App",components:{HelloWorld:f}},m=b,h=(r("034f"),Object(p["a"])(m,a,o,!1,null,null,null)),d=h.exports;n["a"].config.productionTip=!1,new n["a"]({render:function(t){return t(d)}}).$mount("#app")},"85ec":function(t,e,r){},ab29:function(t,e,r){},b132:function(t,e,r){"use strict";r("ab29")},cf05:function(t,e){t.exports="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAFAAAABQCAIAAAABc2X6AAAEjElEQVR4nOybW29bRRDHZ45vOb5f4kttx7k0SlNqqyFNSyp4KPABEIhHviISb4gnBA+8pNAmJCUpgQbZ8hXLcWI78eWcQVFaSsRZO7Z213By/k9J1js7v8zszJ61bc8+2ISbJGXaDsiWBWx2WcBmlwVsdlnAZpcFbHZZwGbXjQO28zWX3oDQArgj6FRBcYGiAGnQ60CnQSdFqB/CcZ4QkO+iY4kzcCKHgdRVHhuoTlCDGFmExfehXaffv6PqC77LjiHZKe2JYO4zJfc5Kg6SvPSlprOHY6v44IvpME+taPlTSvbTKaw+zSodXcHUuuxFp9yWlp6g3SU1sfkD6xp123R2TP3OaBKnGzObUv/pnNvS7lfUPQUavLGu6rPLuPABeiLM3pvegKMfSO9Las62WDLN0dzgHEB/+6s+wFYVitsUzKAaMEayObB7Aqcljl4Mk4x00vu4+6XeY2d4ck3e2UvS/um18dX3TGB/Ej1RSaVLXsEoPodum0mVyEnyRB6wPoDyDhs4C4Aygiy1JRSfM4dm/BhakLGTpQJ36tCqMMN4KyfDB6nATi+5I8zR2CraxD9OSAWe20DFzsxbmxOjq8KzWh6w4qDU+gieRM5EwMn76HCP4AkvgMsn1g1ZwEhz742OHioYz4p1RBLw7Aq6Q9dK10RWbFZLAp6/9idnfHH0JQR6IgPYn6Tg3BgLJUQ2ZBnA84/HWyVxD1GYX8KB1SBF74w3xenF8JKoE4hw4PQjBRWDOqRrVNkbcswU5ZhYYPsMpNaMh46PoLynsybOroBNzOWeWODU+sWB0XCo+pIar0DrG1PZHBi7K6Q/CQRGheYeGjtNRLUD0PrYOJL98CQQOJ5Fl88YuJmHXuvih8ovzOnBDM4E+HslEDizyczJ8ptyVTsgVlYjooiGLAo4tAS+mDGwPqDKi9eQWhfrh0wjIh6eRAHPs8Nb3YfB2dvR0i5zG3si6E9xrtVCgD0xiCwxgQtPrzDUD6l/PqR0cQ6yEOAhu/ekRM3Clb+QhrV9JnD8HiLXt4P4Azu9lHiHOVrYMmAr7zJf71BxiLUJxB84/ZB5cdVtU9noONn4g85PmUHOPEaOV9acgW0OSrMvrkrPiDSjUcLKHtOmN4qpd7ntZM7At9bQoTK6kQaFH5mBKrNrNQAsf6x44zz84/x2KVL2E4UFXN3XS9vMQPVaEL8LTo/xCxQ7JLKAChCBGqbQPLYqE/rIswJG76AaZiLlt0ZML/1Myx8xp9tdePvJ69Feh0rbE+5qnik9pBudlqmZHzG9tEO6xtEdY3EDDqQhmGaH9+nogPRa+PeRU5y4AWfY1869DlXYnfaf+u1bGjCeJXiJD/BMCIZcXBWfgT64lp3uCRx8/X8AzjwCw4ury7urwk9jMJR34OU3OumisDlUafsMJO8z8/nPX6nbHM9gfguOC3T7QwgvXjwV//sFZ41J/LwU/pe/TuvwUjCFrgDYnaD3oXcG3Sa0atDvTG6T8wfT+KrfwtoBZ5s37isAFrDZZQGbXRaw2WUBm10WsNllAZtdFrDZ9VcAAAD//4n9URuJtvpjAAAAAElFTkSuQmCC"}});
//# sourceMappingURL=app.ac4d642d.js.map