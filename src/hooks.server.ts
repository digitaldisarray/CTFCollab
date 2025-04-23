import type { Handle } from '@sveltejs/kit'

// parse the jwt token for username and admin bool
const parseJwt = (token: String) => {
    try {
        return JSON.parse(atob(token.split('.')[1]));
    } catch (e) {
        return null;
    }
};

export const handle: Handle = async ({ event, resolve }) => {
    const jwt = event.cookies.get('token')
   
    if(jwt){
        let resp = parseJwt(jwt)
        
        if(resp != null){
            event.locals.user = {
                username: resp.name,
                isAdmin: resp.isAdmin
            };
        }
        
    } else {
        event.locals.user = null;
    }

    return resolve(event);
}