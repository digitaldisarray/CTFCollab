// See https://svelte.dev/docs/kit/types#app.d.ts
// for information about these interfaces
type User = {
	username: String,
	isAdmin: Boolean,
}
declare global {
	namespace App {
		// interface Error {}
		interface Locals{
			user: User | null;
		}
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}

	}
}

export {};
