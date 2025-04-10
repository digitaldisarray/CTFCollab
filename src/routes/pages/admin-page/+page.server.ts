import type { PageServerLoad, Actions } from "../admin-page/$types.js";
import { fail } from "@sveltejs/kit";
import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";
import { formSchema } from "../admin-page/schema.js";
 
export const load: PageServerLoad = async ({cookies}) => {

    const sessionCookie = cookies.get('session')
    return {
        form: await superValidate(zod(formSchema)),
        cookie: sessionCookie
    };
};
 
export const actions: Actions = {
 default: async (event) => {
  const form = await superValidate(event, zod(formSchema));
  if (!form.valid) {
   return fail(400, {
    form,
   });
  }
  return {
   form,
  };
 },
};