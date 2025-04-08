/** @type {import('./$types').LayoutServerLoad} */
export function load({ locals }){

    if (locals.user) {
		console.log(locals.user.username);
	}
    return {
        user: locals.user && {
            username: locals.user.username,
            isAdmin: locals.user.username
        }
    }
}