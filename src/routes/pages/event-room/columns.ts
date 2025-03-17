import type { ColumnDef } from "@tanstack/table-core";
import { createRawSnippet } from "svelte";
import { renderSnippet } from "$lib/components/ui/data-table/index.js";
import { renderComponent } from "$lib/components/ui/data-table/index.js";
import DataTableActions from "./data-table-actions.svelte";
import { Checkbox } from "$lib/components/ui/checkbox/index.js";
import DataTableNameButton from "./data-name-button.svelte";
import { writable } from "svelte/store";
import {type CTF } from "../admin-page/columns.js"
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


export const challenges = writable<Challenge[]>([])
export const currentCTF = writable<CTF>()

export const columns: ColumnDef<Challenge>[] = [

    {
        accessorKey: "status",
        header: "Status",
          cell: ({ row }) =>
            renderComponent(Checkbox, {
              checked: row.getIsSelected(),
              onCheckedChange: (value) => row.toggleSelected(!!value),
              "aria-label": "Select row",
            }),
          enableSorting: false,
          enableHiding: false,
    },
    {
        accessorKey: "name",
        header: ({ column }) =>
            renderComponent(DataTableNameButton, {
              onclick: () => column.toggleSorting(column.getIsSorted() === "asc"),
            }),
      
    },
    {
        accessorKey: "active_members",
        header: () => {
          const active_membersHeaderSnippet = createRawSnippet(() => ({
            render: () => `<div class="text-right">Active Members</div>`,
          }));
          return renderSnippet(active_membersHeaderSnippet, "");
        },
        
        cell: ({ row }) => {
          const formatter = new Intl.NumberFormat("en-US", {
            style: "decimal",
          });
     
          const active_membersCellSnippet = createRawSnippet<[string]>((getAmount) => {
            const active_members = getAmount();
            return {
              render: () => `<div class="text-right font-medium">${active_members}</div>`,
            };
          });
     
          return renderSnippet(
            active_membersCellSnippet,
            formatter.format(parseFloat(row.getValue("active_members")))
          );
          
          
        },
      },
      {
        id: "actions",
        cell: ({ row }) => {
          // You can pass whatever you need from `row.original` to the component
          return renderComponent(DataTableActions, { 
            id: row.original.id,
            description: row.original.description
          });
        },
      },
    ];

// export const data: Challenge[] = [
//     {
//       id: "1",
//       active_members: 2,
//       status: "pending",
//       name: "Challenge 3",
//     },
//     {
//       id: "2",
//       active_members: 1,
//       status: "processing",
//       name: "Challenge 1",
//     },
//     {
//         id: "3",
//         active_members: 3,
//         status: "processing",
//         name: "Challenge 2",
//     },
//     {
//       id: "4",
//       active_members: 2,
//       status: "pending",
//       name: "Challenge 5",
//     },
//     {
//       id: "5",
//       active_members: 1,
//       status: "processing",
//       name: "Challenge 6",
//     },
//     {
//         id: "6",
//         active_members: 3,
//         status: "processing",
//         name: "Challenge 4",
//     },
// ];