(window.webpackJsonp=window.webpackJsonp||[]).push([[19],{505:function(t,e,o){},529:function(t,e,o){"use strict";o(505)},552:function(t,e,o){"use strict";o.r(e);var s=o(37),r=(o(39),o(40),o(128),o(17),o(65),o(90),o(91),o(25),{props:["value","tree","compact"],data:function(){return{search:{query:null},products:[{label:"sdk",name:"Cosmos<br>SDK",url:"https://docs.cosmos.network/",color:"#5064FB"},{label:"hub",name:"Cosmos<br>Hub",url:"https://hub.cosmos.network/",color:"#BA3FD9"},{label:"ibc",name:"IBC<br>Protocol",url:"https://github.com/cosmos/ics/tree/master/ibc",color:"#E6900A"},{label:"core",name:"Tendermint<br>Core",url:"https://docs.tendermint.com/",color:"#00BB00"}]}},computed:{searchResults:function(){var t=this;return this.$site.pages.filter((function(e){var o=e.headers?e.headers.map((function(t){return t.title})):[],r=e.title;return r&&[r].concat(Object(s.a)(o)).join(" ").toLowerCase().match(t.search.query.toLowerCase())}))},logo:function(){return this.$themeConfig.logo},sidebar:function(){return this.value}}}),i=(o(529),o(0)),a=Object(i.a)(r,(function(){var t=this,e=t.$createElement,o=t._self._c||e;return o("div",{staticStyle:{height:"100%",position:"relative"}},[o("div",{staticClass:"container"},[!0!==t.compact?o("router-link",{staticClass:"logo__container",attrs:{to:"/"}},[o("div",{staticClass:"logo"},[t.$themeConfig.logo&&t.$themeConfig.logo.src?o("div",{staticClass:"logo__img__custom"},[o("img",{attrs:{src:t.$themeConfig.logo.src}})]):o("div",{staticClass:"logo__img"},[o("logo-"+(t.$themeConfig.label||"sdk"),{tag:"component"})],1),t.$themeConfig.logo?t._e():o("div",{staticClass:"logo__text"},[t._v(t._s(t.$site.title||"Documentation"))])])]):t._e(),o("div",{staticClass:"items",class:["footer__compact__"+!(!0!==t.compact)]},[t._l(t.value,(function(e){return o("div",{staticClass:"sidebar",style:{display:0==t.$themeConfig.sidebar.auto&&""===e.title?"none":"block"}},[o("div",{staticClass:"title"},[t._v(t._s(e.title))]),o("client-only",[e.children?o("tm-sidebar-tree",{staticClass:"section",attrs:{value:e.children,tree:t.tree,level:0}}):t._e()],1)],1)})),o("div",{staticClass:"sidebar version"},[o("tm-select-version")],1)],2),t.$themeConfig.custom?t._e():o("div",{staticClass:"footer",class:["footer__compact__"+!(!0!==t.compact)]},t._l(t.products,(function(e){return t.$themeConfig.label!=e.label?o("a",{staticClass:"footer__item",style:{"--color":e.color},attrs:{href:e.url,target:"_blank",rel:"noreferrer noopener"}},[o("tm-logo-"+e.label,{tag:"component",staticClass:"footer__item__icon"}),o("div",{staticClass:"footer__item__title",domProps:{innerHTML:t._s(t.md(e.name))}})],1):t._e()})),0)],1)])}),[],!1,null,"b79cecce",null);e.default=a.exports}}]);