import { f as fail } from './stringify-B14jp4IP.js';
import './client-CZszgVC4.js';
import { z, s as superValidate, a as zod } from './superValidate-C3ZLgWDn.js';

const formSchema = z.object({
  challengeName: z.string().min(2).max(50),
  challengeDetails: z.string().min(0).max(50),
  challengeFlag: z.string().min(0).max(50)
});

const load = async () => {
  return {
    form: await superValidate(zod(formSchema))
  };
};
const actions = {
  default: async (event) => {
    const form = await superValidate(event, zod(formSchema));
    if (!form.valid) {
      return fail(400, {
        form
      });
    }
    return {
      form
    };
  }
};

var _page_server_ts = /*#__PURE__*/Object.freeze({
  __proto__: null,
  actions: actions,
  load: load
});

const index = 5;
let component_cache;
const component = async () => component_cache ??= (await import('./_page.svelte-C_IDe1NP.js')).default;
const server_id = "src/routes/pages/event-room/+page.server.ts";
const imports = ["_app/immutable/nodes/5.D40DolL_.js","_app/immutable/chunks/CoWB_c7W.js","_app/immutable/chunks/Cgeft74Y.js","_app/immutable/chunks/5z2gxC1j.js","_app/immutable/chunks/DJoc1YTT.js","_app/immutable/chunks/BqR-gka9.js","_app/immutable/chunks/tsR4g0rO.js","_app/immutable/chunks/DIS1lw11.js","_app/immutable/chunks/Bii4bBgB.js","_app/immutable/chunks/DOLpMf4M.js","_app/immutable/chunks/DcQ22hbZ.js","_app/immutable/chunks/Iw3L7d7c.js","_app/immutable/chunks/HeQKgpHt.js","_app/immutable/chunks/Br5c-oLs.js","_app/immutable/chunks/DojlfI03.js","_app/immutable/chunks/Cv5cQigG.js","_app/immutable/chunks/CKQFvGxp.js","_app/immutable/chunks/DhcJXaWI.js","_app/immutable/chunks/GU8K1UcO.js","_app/immutable/chunks/C-KXs9MO.js"];
const stylesheets = ["_app/immutable/assets/5.D6rQkR8s.css","_app/immutable/assets/table-row.BwSOqcov.css"];
const fonts = [];

var _5 = /*#__PURE__*/Object.freeze({
  __proto__: null,
  component: component,
  fonts: fonts,
  imports: imports,
  index: index,
  server: _page_server_ts,
  server_id: server_id,
  stylesheets: stylesheets
});

export { _5 as _, formSchema as f };
//# sourceMappingURL=5-biIfybmS.js.map
