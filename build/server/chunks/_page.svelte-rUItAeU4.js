import { V as copy_payload, W as assign_payload, t as pop, p as push } from './index-D3Djb8Yj.js';
import { B as Button } from './button-dEtd2XS0.js';
import { I as Input } from './input-DTwr2lW9.js';
import './client-CZszgVC4.js';
import './exports-wkiByGb4.js';

function _page($$payload, $$props) {
  push();
  let roomcode = "";
  let $$settled = true;
  let $$inner_payload;
  function $$render_inner($$payload2) {
    $$payload2.out += `<div class="welcome-container svelte-128j8gz"><div class="absolute left-4 top-4 md:left-8 md:top-7"><a href="/" class="logo svelte-128j8gz"><span class="ctf svelte-128j8gz">CTF</span> <span class="collab svelte-128j8gz">Collab</span></a></div> <div class="absolute right-4 top-4 md:right-8 md:top-8">`;
    Button($$payload2, {
      href: "/pages/signin",
      variant: "ghost",
      children: ($$payload3) => {
        $$payload3.out += `<!---->Login`;
      },
      $$slots: { default: true }
    });
    $$payload2.out += `<!----> <span class="separator svelte-128j8gz">|</span> `;
    Button($$payload2, {
      href: "/pages/about",
      variant: "ghost",
      children: ($$payload3) => {
        $$payload3.out += `<!---->About`;
      },
      $$slots: { default: true }
    });
    $$payload2.out += `<!----> <span class="separator svelte-128j8gz">|</span> `;
    Button($$payload2, {
      href: "/pages/event-room",
      variant: "ghost",
      children: ($$payload3) => {
        $$payload3.out += `<!---->Test Event Room`;
      },
      $$slots: { default: true }
    });
    $$payload2.out += `<!----></div> <header class="svelte-128j8gz"><h1 class="svelte-128j8gz">Welcome to CTF-Collab!</h1> <p class="svelte-128j8gz">To join an event, enter room code below or login with existing account</p> <div class="form-container svelte-128j8gz"><form class="flex w-full max-w-sm items-center space-x-2">`;
    Input($$payload2, {
      type: "room-code",
      placeholder: "enter room-code here...",
      get value() {
        return roomcode;
      },
      set value($$value) {
        roomcode = $$value;
        $$settled = false;
      }
    });
    $$payload2.out += `<!----> `;
    Button($$payload2, {
      type: "submit",
      children: ($$payload3) => {
        $$payload3.out += `<!---->Submit`;
      },
      $$slots: { default: true }
    });
    $$payload2.out += `<!----></form></div> `;
    {
      $$payload2.out += "<!--[!-->";
    }
    $$payload2.out += `<!--]--></header> <main class="svelte-128j8gz"></main> <footer class="svelte-128j8gz"><p>© 2025 CTF-Collab. All rights reserved.</p></footer></div>`;
  }
  do {
    $$settled = true;
    $$inner_payload = copy_payload($$payload);
    $$render_inner($$inner_payload);
  } while (!$$settled);
  assign_payload($$payload, $$inner_payload);
  pop();
}

export { _page as default };
//# sourceMappingURL=_page.svelte-rUItAeU4.js.map
