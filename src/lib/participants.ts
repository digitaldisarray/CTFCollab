import { writable } from "svelte/store";

export type Participant = {
  id: number;
  ctf_id: number;
  user_id: { Int32: number, Valid: boolean };
  guest_id: { Int32: number, Valid: boolean };
  username: { String: string, Valid: boolean };
  nickname: { String: string, Valid: boolean };
};

export const participants = writable<Participant[]>([]);
