import type { PageServerLoad, Actions } from "../admin-page/$types.js";
import { fail } from "@sveltejs/kit";
import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";
import { changePasswordForm, formSchema } from "../admin-page/schema.js";
 
export const load: PageServerLoad = async ({cookies}) => {

    
    return {
        form: await superValidate(zod(formSchema)),
        changePasswordForm: await superValidate(zod(changePasswordForm))
        
    };
};
 
export const actions: Actions = {
    default: async (event) => {
        const form = await superValidate(event, zod(formSchema));
        const changePassForm = await superValidate(event, zod(changePasswordForm));
        if (!form.valid) {
            return fail(400, {
                    form,
                    changePassForm
                });
        }
        return {
            form,
            changePassForm
        };
    },
};