(self.webpackChunk=self.webpackChunk||[]).push([[61],{7108:(n,t,e)=>{"use strict";const i={data:function(){return{menu:window.init.menu,cruds:window.init.cruds,permissions:window.init.permissions.permissions,extra:window.init.permissions.extra}},methods:{can:function(n){return!!this.permissions[n.id]&&!!this.permissions[n.id].show},getTitle:function(n){if("crud"==n.link_to){var t=this.cruds.findIndex((function(t){return t.id==n.url}));return t>=0?this.cruds[t].title:""}return n.title}}};var r={"page-nav":(0,e(51900).Z)(i,(function(){var n=this,t=n.$createElement,e=n._self._c||t;return e("ul",n._l(n.menu,(function(t,i){return n.can(t)?e("li",{key:i},["link"!=t.link_to&&"router-link"!=t.link_to?e("router-link",{attrs:{to:"/p/"+t.id}},[t.icon?e("i",{class:t.icon}):n._e(),n._v(" "),e("span",{domProps:{innerHTML:n._s(n.getTitle(t))}})]):n._e(),n._v(" "),"router-link"==t.link_to?e("router-link",{attrs:{to:t.url}},[t.icon?e("i",{class:t.icon}):n._e(),n._v(" "),e("span",{domProps:{innerHTML:n._s(n.getTitle(t))}})]):n._e(),n._v(" "),"link"==t.link_to?e("a",{attrs:{href:t.url,target:"_blank"}},[t.icon?e("i",{class:t.icon}):n._e(),n._v(" "),e("span",{domProps:{innerHTML:n._s(n.getTitle(t))}})]):n._e()],1):n._e()})),0)}),[],!1,null,null,null).exports},o=function n(t,e){n.installed||Object.keys(r).forEach((function(n){return t.component(n,r[n])}))};"undefined"!=typeof window&&window.Vue&&o(window.Vue)},51900:(n,t,e)=>{"use strict";function i(n,t,e,i,r,o,s,u){var c,l="function"==typeof n?n.options:n;if(t&&(l.render=t,l.staticRenderFns=e,l._compiled=!0),i&&(l.functional=!0),o&&(l._scopeId="data-v-"+o),s?(c=function(n){(n=n||this.$vnode&&this.$vnode.ssrContext||this.parent&&this.parent.$vnode&&this.parent.$vnode.ssrContext)||"undefined"==typeof __VUE_SSR_CONTEXT__||(n=__VUE_SSR_CONTEXT__),r&&r.call(this,n),n&&n._registeredComponents&&n._registeredComponents.add(s)},l._ssrRegister=c):r&&(c=u?function(){r.call(this,(l.functional?this.parent:this).$root.$options.shadowRoot)}:r),c)if(l.functional){l._injectStyles=c;var a=l.render;l.render=function(n,t){return c.call(t),a(n,t)}}else{var d=l.beforeCreate;l.beforeCreate=d?[].concat(d,c):[c]}return{exports:n,options:l}}e.d(t,{Z:()=>i})}},0,[[7108,245]]]);
//# sourceMappingURL=page-nav.js.map