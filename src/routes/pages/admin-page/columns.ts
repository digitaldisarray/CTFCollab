import type { ColumnDef } from "@tanstack/table-core";
import { createRawSnippet } from "svelte";
import { renderSnippet } from "$lib/components/ui/data-table/index.js";
import { renderComponent } from "$lib/components/ui/data-table/index.js";
import DataTableActions from "./data-table-actions.svelte";
import { Checkbox } from "$lib/components/ui/checkbox/index.js";
import DataTableNameButton from "./data-name-button.svelte";
import DataTableStatusButton from "./data-status-button.svelte";
import DataTableDateButton from "./data-date-button.svelte";
import { writable } from "svelte/store";
import { get } from "svelte/store";
import { participants } from "$lib/participants";

/**
 *CURRENT USE: 
  id = phrase
  members = default to 0
  status = default to pending
  date = start_date 
  name = name
 */
export type Challenge = {
    id: string;
    members: number;
    status: "pending" | "completed" | "in progress" | "canceled";
    date: string;
    name: string;
};


export type CTF = {
    ctf_id: number,
    ctf_author_id: string | null;  
    ctf_description: string | null;
    ctf_name: string;
    start_date: Date;
    end_date: Date;
    phrase: string;
}

export const ctfData = writable<Challenge[]>([]) // so the CTF list is rememeber across states
const memberCount = get(participants).length;

// format CTF's to challenges
export function formatData(ctf: Array<CTF>): Challenge[] {
  return ctf.map((c) => {
    const today = new Date();
    const ctfDate = new Date(c.start_date);
    let status: Challenge["status"] = "pending";

    if (ctfDate < today) {
      status = "completed";
    } else if (ctfDate.toDateString() === today.toDateString()) {
      status = "in progress";
    }

    return {
      id: String(c.phrase),
      name: c.ctf_name,
      date: `${formatDateOnly(c.start_date)} - ${formatDateOnly(c.end_date)}`,
      status,
      members: memberCount,
    };
  });
}

export function formatDateOnly(date: Date | string): string {
  const d = typeof date === "string" ? new Date(date) : date;
  return d.toISOString().split("T")[0]; // "2025-06-17"
}


export const columns: ColumnDef<Challenge>[] = [

    {
      accessorKey: "status",
      header: ({ column }) =>
        renderComponent(DataTableStatusButton, {
          onclick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        }),
      cell: ({ row }) => {
        const status = row.getValue("status") as Challenge["status"];
        const statusColor = {
          pending: "text-yellow-500",
          "in progress": "text-orange-500",
          completed: "text-green-500",
          canceled: "text-red-500",
        }[status];

        const statusSnippet = createRawSnippet<[string]>((getStatus) => {
          const status = getStatus();
          return {
            render: () => `<div class="font-medium ${statusColor}">${status}</div>`,
          };
        });

        return renderSnippet(statusSnippet, status);
      },
    },
    {
        accessorKey: "name",
        header: ({ column }) =>
            renderComponent(DataTableNameButton, {
              onclick: () => column.toggleSorting(column.getIsSorted() === "asc"),
            }),
     
    },
    {
      accessorKey: "date",
      header: ({ column }) =>
          renderComponent(DataTableDateButton, {
            onclick: () => column.toggleSorting(column.getIsSorted() === "asc"),
          }),
   
    },
    {
        accessorKey: "members",
        header: () => {
          const membersHeaderSnippet = createRawSnippet(() => ({
            render: () => `<div class="text-right">Members</div>`,
          }));
          return renderSnippet(membersHeaderSnippet, "");
        },
       
        cell: ({ row }) => {
          const formatter = new Intl.NumberFormat("en-US", {
            style: "decimal",
          });
     
          const membersCellSnippet = createRawSnippet<[string]>((getAmount) => {
            const members = getAmount();
            return {
              render: () => `<div class="text-right font-medium">${members}</div>`,
            };
          });
     
          return renderSnippet(
            membersCellSnippet,
            formatter.format(parseFloat(row.getValue("members")))
          );
         
         
        },
      },
      {
        id: "actions",
        cell: ({ row }) => {
          // You can pass whatever you need from `row.original` to the component
          return renderComponent(DataTableActions, { id: row.original.id });
        },
      },
    ];
