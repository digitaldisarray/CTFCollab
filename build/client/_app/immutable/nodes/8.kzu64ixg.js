import{a as g,t as _,e as T}from"../chunks/CoWB_c7W.js";import"../chunks/GU8K1UcO.js";import{u as C,w as D,P as c,Q as s,D as a,R as h,S as y,n as i,T as o}from"../chunks/Cgeft74Y.js";import{i as E}from"../chunks/C-KXs9MO.js";import{B as I}from"../chunks/DcQ22hbZ.js";import{I as k}from"../chunks/CKQFvGxp.js";import{g as P}from"../chunks/Cv5cQigG.js";var S=_('<div class="login-container svelte-4k3tpc"><header class="svelte-4k3tpc"><h1 class="svelte-4k3tpc">Login as Administrator</h1> <p class="svelte-4k3tpc">Enter your email below to access existing admin account</p> <div class="form-container svelte-4k3tpc"><form class="flex flex-col w-full max-w-sm items-center space-y-2"><!> <!> <!></form> <p class="svelte-4k3tpc"></p></div></header> <main class="svelte-4k3tpc"></main> <footer class="svelte-4k3tpc"><p>&copy; 2025 CTF-Collab. All rights reserved.</p></footer></div>');function N(w,x){C(x,!1);let r=y(""),l=y("");const $=async e=>{e.preventDefault();const f={username:a(r),password:a(l)};try{(await fetch("http://localhost:1337/users/login",{method:"POST",headers:{"Content-Type":"application/json"},credentials:"include",body:JSON.stringify(f)})).ok?P("/pages/admin-page"):console.log("Failure")}catch(t){console.error("Error occured",t)}};E();var n=S(),p=s(n),m=c(s(p),4),d=s(m),v=s(d);k(v,{type:"email",placeholder:"me@example.com...",get value(){return a(r)},set value(e){h(r,e)},$$legacy:!0});var u=c(v,2);k(u,{type:"password",placeholder:"Enter your password",get value(){return a(l)},set value(e){h(l,e)},$$legacy:!0});var b=c(u,2);I(b,{href:"src\\routes\\admin-account.svelte",type:"submit",onclick:$,children:(e,f)=>{i();var t=T("Log In");g(e,t)},$$slots:{default:!0}}),o(d),i(2),o(m),o(p),i(4),o(n),g(w,n),D()}export{N as component};
