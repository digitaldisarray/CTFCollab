// See https://svelte.dev/docs/kit/types#app.d.ts
// for information about these interfaces

declare global {
	type User = {
		username: string,
		isAdmin: Boolean,
	}
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
