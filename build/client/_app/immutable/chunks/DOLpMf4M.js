import{a4 as O,a5 as $,a6 as z,a7 as w,a8 as H,R as b,a9 as E,a0 as v,D as h,v as L,aa as V,ab as J,ac as Q,ad as W,Y as X,A as K,ae as k,af as ee,ag as re,ah as G,ai as ne,aj as ie,ak as se,al as B,am as te,M as ae,an as fe,ao as ue,ap as Z,G as C,aq as le,a1 as ce,ar as de,as as A,at as M}from"./Cgeft74Y.js";import{c as oe}from"./tsR4g0rO.js";function I(e,r=null,t){if(typeof e!="object"||e===null||O in e)return e;const f=J(e);if(f!==$&&f!==z)return e;var i=new Map,c=Q(e),P=w(0);c&&i.set("length",w(e.length));var y;return new Proxy(e,{defineProperty(l,n,s){(!("value"in s)||s.configurable===!1||s.enumerable===!1||s.writable===!1)&&H();var u=i.get(n);return u===void 0?(u=w(s.value),i.set(n,u)):b(u,I(s.value,y)),!0},deleteProperty(l,n){var s=i.get(n);if(s===void 0)n in l&&i.set(n,w(v));else{if(c&&typeof n=="string"){var u=i.get("length"),a=Number(n);Number.isInteger(a)&&a<u.v&&b(u,a)}b(s,v),Y(P)}return!0},get(l,n,s){var _;if(n===O)return e;var u=i.get(n),a=n in l;if(u===void 0&&(!a||(_=E(l,n))!=null&&_.writable)&&(u=w(I(a?l[n]:v,y)),i.set(n,u)),u!==void 0){var d=h(u);return d===v?void 0:d}return Reflect.get(l,n,s)},getOwnPropertyDescriptor(l,n){var s=Reflect.getOwnPropertyDescriptor(l,n);if(s&&"value"in s){var u=i.get(n);u&&(s.value=h(u))}else if(s===void 0){var a=i.get(n),d=a==null?void 0:a.v;if(a!==void 0&&d!==v)return{enumerable:!0,configurable:!0,value:d,writable:!0}}return s},has(l,n){var d;if(n===O)return!0;var s=i.get(n),u=s!==void 0&&s.v!==v||Reflect.has(l,n);if(s!==void 0||L!==null&&(!u||(d=E(l,n))!=null&&d.writable)){s===void 0&&(s=w(u?I(l[n],y):v),i.set(n,s));var a=h(s);if(a===v)return!1}return u},set(l,n,s,u){var x;var a=i.get(n),d=n in l;if(c&&n==="length")for(var _=s;_<a.v;_+=1){var g=i.get(_+"");g!==void 0?b(g,v):_ in l&&(g=w(v),i.set(_+"",g))}a===void 0?(!d||(x=E(l,n))!=null&&x.writable)&&(a=w(void 0),b(a,I(s,y)),i.set(n,a)):(d=a.v!==v,b(a,I(s,y)));var p=Reflect.getOwnPropertyDescriptor(l,n);if(p!=null&&p.set&&p.set.call(u,s),!d){if(c&&typeof n=="string"){var S=i.get("length"),m=Number(n);Number.isInteger(m)&&m>=S.v&&b(S,m+1)}Y(P)}return!0},ownKeys(l){h(P);var n=Reflect.ownKeys(l).filter(a=>{var d=i.get(a);return d===void 0||d.v!==v});for(var[s,u]of i)u.v!==v&&!(s in l)&&n.push(s);return n},setPrototypeOf(){V()}})}function Y(e,r=1){b(e,e.v+r)}function F(e,r){return e===r||(e==null?void 0:e[O])===r}function Pe(e={},r,t,f){return W(()=>{var i,c;return X(()=>{i=c,c=[],K(()=>{e!==t(...c)&&(r(e,...c),i&&F(t(...i),e)&&r(null,...i))})}),()=>{k(()=>{c&&F(t(...c),e)&&r(null,...c)})}}),e}const ve={get(e,r){if(!e.exclude.includes(r))return e.props[r]},set(e,r){return!1},getOwnPropertyDescriptor(e,r){if(!e.exclude.includes(r)&&r in e.props)return{enumerable:!0,configurable:!0,value:e.props[r]}},has(e,r){return e.exclude.includes(r)?!1:r in e.props},ownKeys(e){return Reflect.ownKeys(e.props).filter(r=>!e.exclude.includes(r))}};function ye(e,r,t){return new Proxy({props:e,exclude:r},ve)}const _e={get(e,r){if(!e.exclude.includes(r))return h(e.version),r in e.special?e.special[r]():e.props[r]},set(e,r,t){return r in e.special||(e.special[r]=he({get[r](){return e.props[r]}},r,G)),e.special[r](t),M(e.version),!0},getOwnPropertyDescriptor(e,r){if(!e.exclude.includes(r)&&r in e.props)return{enumerable:!0,configurable:!0,value:e.props[r]}},deleteProperty(e,r){return e.exclude.includes(r)||(e.exclude.push(r),M(e.version)),!0},has(e,r){return e.exclude.includes(r)?!1:r in e.props},ownKeys(e){return Reflect.ownKeys(e.props).filter(r=>!e.exclude.includes(r))}};function ge(e,r){return new Proxy({props:e,exclude:r,special:{},version:w(0)},_e)}const pe={get(e,r){let t=e.props.length;for(;t--;){let f=e.props[t];if(A(f)&&(f=f()),typeof f=="object"&&f!==null&&r in f)return f[r]}},set(e,r,t){let f=e.props.length;for(;f--;){let i=e.props[f];A(i)&&(i=i());const c=E(i,r);if(c&&c.set)return c.set(t),!0}return!1},getOwnPropertyDescriptor(e,r){let t=e.props.length;for(;t--;){let f=e.props[t];if(A(f)&&(f=f()),typeof f=="object"&&f!==null&&r in f){const i=E(f,r);return i&&!i.configurable&&(i.configurable=!0),i}}},has(e,r){if(r===O||r===Z)return!1;for(let t of e.props)if(A(t)&&(t=t()),t!=null&&r in t)return!0;return!1},ownKeys(e){const r=[];for(let t of e.props){A(t)&&(t=t());for(const f in t)r.includes(f)||r.push(f)}return r}};function xe(...e){return new Proxy({props:e},pe)}function U(e){for(var r=L,t=L;r!==null&&!(r.f&(ie|se));)r=r.parent;try{return B(r),e()}finally{B(t)}}function he(e,r,t,f){var j;var i=(t&te)!==0,c=!ae||(t&fe)!==0,P=(t&ue)!==0,y=(t&de)!==0,l=!1,n;P?[n,l]=oe(()=>e[r]):n=e[r];var s=O in e||Z in e,u=P&&(((j=E(e,r))==null?void 0:j.set)??(s&&r in e&&(o=>e[r]=o)))||void 0,a=f,d=!0,_=!1,g=()=>(_=!0,d&&(d=!1,y?a=K(f):a=f),a);n===void 0&&f!==void 0&&(u&&c&&ee(),n=g(),u&&u(n));var p;if(c)p=()=>{var o=e[r];return o===void 0?g():(d=!0,_=!1,o)};else{var S=U(()=>(i?C:le)(()=>e[r]));S.f|=re,p=()=>{var o=h(S);return o!==void 0&&(a=void 0),o===void 0?a:o}}if(!(t&G))return p;if(u){var m=e.$$legacy;return function(o,R){return arguments.length>0?((!c||!R||m||l)&&u(R?p():o),o):p()}}var x=!1,q=!1,N=ce(n),D=U(()=>C(()=>{var o=p(),R=h(N);return x?(x=!1,q=!0,R):(q=!1,N.v=o)}));return i||(D.equals=ne),function(o,R){if(arguments.length>0){const T=R?h(D):c&&P?I(o):o;return D.equals(T)||(x=!0,b(N,T),_&&a!==void 0&&(a=T),K(()=>h(D))),o}return h(D)}}export{I as a,Pe as b,ge as l,he as p,ye as r,xe as s};
