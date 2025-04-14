import type { Handle } from '@sveltejs/kit'

export const handle: Handle = async ({ event, resolve }) => {
    const jwt = event.cookies.get('token')
   
    if(jwt){
        let response = await fetch("http://localhost:1337/users/token", {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                Cookie: `token=${jwt}` // hooks.server.ts is beyond the context of the browser, so need to specify the cookie here
            },
        });
  
        if(response.ok){
            let resp = await response.json();
            event.locals.user = {
                username: resp.username,
                isAdmin: resp.isAdmin
            };
            

        } else {
            event.locals.user = null;
        }
        
    } else {
        event.locals.user = null;
    }

    return resolve(event);
}