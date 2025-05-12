import { writable } from "svelte/store";
import type { CTF } from "../../admin-page/columns";
/**
 * CURRENT USE:
 * id = default to nothing
 * active_members = default to 0
 * status = pending
 * name = name
 */
export type Challenge = {
    id: string;
    hedgedoc_url: string;
    active_members: number;
    status: "pending" | "processing" | "success" | "failed";
    name: string;
    description: string;
};

export type CTFChallenge = {
    id: number;
    ctf_id: number;
    name: string;
    description: string;
    flag: string;
    hedgedoc_url: string; 
}

export function formatData(ctfch: Array<CTFChallenge>): Challenge[]{
  let newData = ctfch.map((c) => {
      return {
        id: c.id.toString(),
        hedgedoc_url: c.hedgedoc_url,
        active_members: 0,
        status: "pending" as "pending" | "processing" | "success" | "failed",
        name: c.name,
        description: c.description
      }
  });
  return newData;
}


export const challenges = writable<Challenge[]>([]) //so the challenge list is rememeber across states
export const currentCTF = writable<CTF>()
