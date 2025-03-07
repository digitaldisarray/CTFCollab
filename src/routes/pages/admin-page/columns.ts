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

 
export type Challenge = {
    id: string;
    members: number;
    status: "pending" | "completed" | "in progress" | "canceled";
    date: string;
    name: string;
};


export type CTF = {
    ctf_author_id: string;  
    ctf_description: string;
    ctf_name: string;
    start_date: Date;
    end_date: Date;
    phrase: string;
}

export const columns: ColumnDef<Challenge>[] = [

    {
        accessorKey: "status",
        header: ({ column }) =>
          renderComponent(DataTableStatusButton, {
            onclick: () => column.toggleSorting(column.getIsSorted() === "asc"),
          }),
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


export const ctfData = writable<Challenge[]>([])


export function formatData(ctf: Array<CTF>): Challenge[]{
  let newData = ctf.map((c) => {
      const today = new Date();
      let ctfDate = new Date(c.start_date);
      let status: "completed" | "pending" | "in progress" = "pending";
      
      if(ctfDate < today){ // TODO: add a status to backend 
        status = "completed"
      } else if (ctfDate.toDateString() === today.toDateString()){
        status = "in progress"
      }
      
      return {
        id: String(c.phrase),
        name: c.ctf_name,
        date: ctfDate.toISOString(),
        status: status as "completed" | "pending" | "in progress" | "canceled",
        members: 0,
      }
  });
  return newData;
}

export const data: Challenge[] = [
    {
      id: "1",
      members: 2,
      status: "pending",
      date: "2023-10-01",
      name: "Event 3",
    },
    {
      id: "2",
      members: 1,
      status: "completed",
      date: "2023-10-01",
      name: "Event 1",
    },
    {
        id: "3",
        members: 3,
        status: "completed",
        date: "2023-10-01",
        name: "Event 2",
    },
    {
      id: "4",
      members: 2,
      status: "completed",
      date: "2023-10-01",
      name: "Event 5",
    },
    {
      id: "5",
      members: 1,
      status: "in progress",
      date: "2023-10-01",
      name: "Event 6",
    },
    {
        id: "6",
        members: 3,
        status: "pending",
        date: "2023-10-01",
        name: "Event 4",
    },
];