import{h,ab as j,au as C,av as E,aw as I}from"./Cgeft74Y.js";import{b as R,c as S,d as V,f as B,n as G,g as M,j as P}from"./DJoc1YTT.js";function N(s){var r,i,t="";if(typeof s=="string"||typeof s=="number")t+=s;else if(typeof s=="object")if(Array.isArray(s)){var o=s.length;for(r=0;r<o;r++)s[r]&&(i=N(s[r]))&&(t&&(t+=" "),t+=i)}else for(i in s)s[i]&&(t&&(t+=" "),t+=i);return t}function q(){for(var s,r,i=0,t="",o=arguments.length;i<o;i++)(s=arguments[i])&&(r=N(s))&&(t&&(t+=" "),t+=r);return t}function z(s){return typeof s=="object"?q(s):s??""}function H(s){if(h){var r=!1,i=()=>{if(!r){if(r=!0,s.hasAttribute("value")){var t=s.value;n(s,"value",null),s.value=t}if(s.hasAttribute("checked")){var o=s.checked;n(s,"checked",null),s.checked=o}}};s.__on_r=i,C(i),M()}}function J(s,r){var i=s.__attributes??(s.__attributes={});i.value===(i.value=r??void 0)||s.value===r&&(r!==0||s.nodeName!=="PROGRESS")||(s.value=r??"")}function D(s,r){r?s.hasAttribute("selected")||s.setAttribute("selected",""):s.removeAttribute("selected")}function n(s,r,i,t){var o=s.__attributes??(s.__attributes={});h&&(o[r]=s.getAttribute(r),r==="src"||r==="srcset"||r==="href"&&s.nodeName==="LINK")||o[r]!==(o[r]=i)&&(r==="style"&&"__styles"in s&&(s.__styles={}),r==="loading"&&(s[E]=i),i==null?s.removeAttribute(r):typeof i!="string"&&$(s).includes(r)?s[r]=i:s.setAttribute(r,i))}function Q(s,r,i,t,o=!1,v=!1,K=!1){var l=r||{},y=s.tagName==="OPTION";for(var p in r)p in i||(i[p]=null);i.class&&(i.class=z(i.class));var w=$(s),O=s.__attributes??(s.__attributes={});for(const a in i){let f=i[a];if(y&&a==="value"&&f==null){s.value=s.__value="",l[a]=f;continue}var g=l[a];if(f!==g){l[a]=f;var A=a[0]+a[1];if(A!=="$$"){if(A==="on"){const e={},_="$$"+a;let c=a.slice(2);var d=P(c);if(R(c)&&(c=c.slice(0,-7),e.capture=!0),!d&&g){if(f!=null)continue;s.removeEventListener(c,l[_],e),l[_]=null}if(f!=null)if(d)s[`__${c}`]=f,V([c]);else{let L=function(T){l[a].call(this,T)};l[_]=S(c,s,L,e)}else d&&(s[`__${c}`]=void 0)}else if(a==="style"&&f!=null)s.style.cssText=f+"";else if(a==="autofocus")B(s,!!f);else if(a==="__value"||a==="value"&&f!=null)s.value=s[a]=s.__value=f;else if(a==="selected"&&y)D(s,f);else{var u=a;o||(u=G(u));var b=u==="defaultValue"||u==="defaultChecked";if(f==null&&!v&&!b)if(O[a]=null,u==="value"||u==="checked"){let e=s;if(u==="value"){let _=e.defaultValue;e.removeAttribute(u),e.defaultValue=_}else{let _=e.defaultChecked;e.removeAttribute(u),e.defaultChecked=_}}else s.removeAttribute(a);else b||w.includes(u)&&(v||typeof f!="string")?s[u]=f:typeof f!="function"&&(h&&(u==="src"||u==="href"||u==="srcset")||n(s,u,f))}a==="style"&&"__styles"in s&&(s.__styles={})}}}return l}var k=new Map;function $(s){var r=k.get(s.nodeName);if(r)return r;k.set(s.nodeName,r=[]);for(var i,t=s,o=Element.prototype;o!==t;){i=I(t);for(var v in i)i[v].set&&r.push(v);t=j(t)}return r}export{Q as a,J as b,q as c,H as r,n as s};
