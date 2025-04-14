import { X as store_get, Y as unsubscribe_stores, t as pop, p as push } from './index-D3Djb8Yj.js';
import { p as page } from './stores-BfAkj-Gl.js';
import './client-CZszgVC4.js';
import './exports-wkiByGb4.js';

function _page($$payload, $$props) {
  push();
  var $$store_subs;
  new URLSearchParams(store_get($$store_subs ??= {}, "$page", page).url.search).get("code");
  new URLSearchParams(store_get($$store_subs ??= {}, "$page", page).url.search).get("challenge") || "";
  $$payload.out += `<header class="logo-header svelte-15fx9b1"><div class="absolute left-4 top-4 md:left-8 md:top-7"><a href="/" class="logo svelte-15fx9b1"><span class="ctf svelte-15fx9b1">CTF</span> <span class="collab svelte-15fx9b1">Collab</span></a></div></header> <div class="iframe-container svelte-15fx9b1">`;
  {
    $$payload.out += "<!--[!-->";
    $$payload.out += `<p>Loading challenge...</p>`;
  }
  $$payload.out += `<!--]--></div>`;
  if ($$store_subs) unsubscribe_stores($$store_subs);
  pop();
}

export { _page as default };
//# sourceMappingURL=_page.svelte-1fgL6kBa.js.map
