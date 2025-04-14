import { V as copy_payload, W as assign_payload, t as pop, p as push } from './index-D3Djb8Yj.js';
import { B as Button } from './button-dEtd2XS0.js';
import { I as Input } from './input-DTwr2lW9.js';
import { g as goto } from './client-CZszgVC4.js';
import './exports-wkiByGb4.js';

function _page($$payload, $$props) {
  push();
  let email = "";
  let password = "";
  const login = async (e) => {
    e.preventDefault();
    const loginData = { "username": email, "password": password };
    try {
      const response = await fetch("http://localhost:1337/users/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        // Tell fetch to include cookies
        credentials: "include",
        body: JSON.stringify(loginData)
      });
      if (response.ok) {
        goto("/pages/admin-page");
      } else {
        console.log("Failure");
      }
    } catch (error) {
      console.error("Error occured", error);
    }
  };
  let $$settled = true;
  let $$inner_payload;
  function $$render_inner($$payload2) {
    $$payload2.out += `<div class="login-container svelte-4k3tpc"><header class="svelte-4k3tpc"><h1 class="svelte-4k3tpc">Login as Administrator</h1> <p class="svelte-4k3tpc">Enter your email below to access existing admin account</p> <div class="form-container svelte-4k3tpc"><form class="flex flex-col w-full max-w-sm items-center space-y-2">`;
    Input($$payload2, {
      type: "email",
      placeholder: "me@example.com...",
      get value() {
        return email;
      },
      set value($$value) {
        email = $$value;
        $$settled = false;
      }
    });
    $$payload2.out += `<!----> `;
    Input($$payload2, {
      type: "password",
      placeholder: "Enter your password",
      get value() {
        return password;
      },
      set value($$value) {
        password = $$value;
        $$settled = false;
      }
    });
    $$payload2.out += `<!----> `;
    Button($$payload2, {
      href: "src\\routes\\admin-account.svelte",
      type: "submit",
      onclick: login,
      children: ($$payload3) => {
        $$payload3.out += `<!---->Log In`;
      },
      $$slots: { default: true }
    });
    $$payload2.out += `<!----></form> <p class="svelte-4k3tpc"></p></div></header> <main class="svelte-4k3tpc"></main> <footer class="svelte-4k3tpc"><p>© 2025 CTF-Collab. All rights reserved.</p></footer></div>`;
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
//# sourceMappingURL=_page.svelte-DAwAqoWp.js.map
