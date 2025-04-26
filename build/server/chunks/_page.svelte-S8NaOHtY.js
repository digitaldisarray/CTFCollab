import { Z as bind_props, p as push, V as copy_payload, W as assign_payload, t as pop, _ as spread_props, X as store_get, a6 as store_mutate, Y as unsubscribe_stores } from './index-D3Djb8Yj.js';
import { R as Root, T as Trigger, P as Popover_content, F as Form_field, a as Form_button, c as Control, d as Form_description, e as Form_field_errors, v as Form_label } from './index3-VeECvYOq.js';
import { I as Input } from './input-DTwr2lW9.js';
import { f as formSchema } from './7-CJRPEzpW.js';
import './client-CZszgVC4.js';
import { b as superForm, c as zodClient } from './superValidate-C3ZLgWDn.js';
import './stringify-B14jp4IP.js';
import './button-dEtd2XS0.js';
import './events-DkGbpcOb.js';
import './exports-wkiByGb4.js';
import './stores-BfAkj-Gl.js';

function Settings_form($$payload, $$props) {
  push();
  var $$store_subs;
  let { data } = $$props;
  const form = superForm(data.form, { validators: zodClient(formSchema) });
  const { form: formData, enhance } = form;
  let $$settled = true;
  let $$inner_payload;
  function $$render_inner($$payload2) {
    $$payload2.out += `<form method="POST"><!---->`;
    Form_field($$payload2, {
      form,
      name: "username",
      children: ($$payload3) => {
        $$payload3.out += `<!---->`;
        {
          let children = function($$payload4, { props }) {
            $$payload4.out += `<!---->`;
            Form_label($$payload4, {
              children: ($$payload5) => {
                $$payload5.out += `<!---->Username`;
              },
              $$slots: { default: true }
            });
            $$payload4.out += `<!----> `;
            Input($$payload4, spread_props([
              props,
              {
                get value() {
                  return store_get($$store_subs ??= {}, "$formData", formData).username;
                },
                set value($$value) {
                  store_mutate($$store_subs ??= {}, "$formData", formData, store_get($$store_subs ??= {}, "$formData", formData).username = $$value);
                  $$settled = false;
                }
              }
            ]));
            $$payload4.out += `<!---->`;
          };
          Control($$payload3, { children, $$slots: { default: true } });
        }
        $$payload3.out += `<!----> <!---->`;
        Form_description($$payload3, {
          children: ($$payload4) => {
            $$payload4.out += `<!---->This is your public display name.`;
          },
          $$slots: { default: true }
        });
        $$payload3.out += `<!----> <!---->`;
        Form_field_errors($$payload3, {});
        $$payload3.out += `<!---->`;
      },
      $$slots: { default: true }
    });
    $$payload2.out += `<!----> <!---->`;
    Form_button($$payload2, {
      children: ($$payload3) => {
        $$payload3.out += `<!---->Submit`;
      },
      $$slots: { default: true }
    });
    $$payload2.out += `<!----></form>`;
  }
  do {
    $$settled = true;
    $$inner_payload = copy_payload($$payload);
    $$render_inner($$inner_payload);
  } while (!$$settled);
  assign_payload($$payload, $$inner_payload);
  if ($$store_subs) unsubscribe_stores($$store_subs);
  pop();
}
function _page($$payload, $$props) {
  let data = $$props["data"];
  $$payload.out += `<div class="about svelte-13msih2"><h1 class="svelte-13msih2">New Event Placeholder</h1> <p>Event options will go here</p> `;
  Root($$payload, {
    children: ($$payload2) => {
      Trigger($$payload2, {
        children: ($$payload3) => {
          $$payload3.out += `<!---->Open`;
        },
        $$slots: { default: true }
      });
      $$payload2.out += `<!----> `;
      Popover_content($$payload2, {
        children: ($$payload3) => {
          Settings_form($$payload3, { data });
        },
        $$slots: { default: true }
      });
      $$payload2.out += `<!---->`;
    },
    $$slots: { default: true }
  });
  $$payload.out += `<!----> <a href="/">Home</a></div>`;
  bind_props($$props, { data });
}

export { _page as default };
//# sourceMappingURL=_page.svelte-S8NaOHtY.js.map
