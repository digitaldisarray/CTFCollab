<script lang="ts" generics="TData, TValue">
    import { Button } from "$lib/components/ui/button/index.js";
    import { Input } from "$lib/components/ui/input/index.js";
    import { type Challenge } from "./columns.js"
    import { goto } from '$app/navigation';

    import {
        type ColumnDef,
        type PaginationState,
        type RowSelectionState,
        type SortingState,
        type ColumnFiltersState,
        getCoreRowModel,
        getPaginationRowModel,
        getSortedRowModel,
        getFilteredRowModel,
       
    } from "@tanstack/table-core";


    import {
     createSvelteTable,
     FlexRender,
    } from "$lib/components/ui/data-table/index.js";
    import * as Table from "$lib/components/ui/table/index.js";
   
    type DataTableProps<TData, TValue> = {
    data: TData[];
    columns: ColumnDef<TData, TValue>[];
    };
   
    let { data, columns }: DataTableProps<TData, TValue> = $props();


    let pagination = $state<PaginationState>({ pageIndex: 0, pageSize: 10 });
    let sorting = $state<SortingState>([]);
    let columnFilters = $state<ColumnFiltersState>([]);
   
    const table = createSvelteTable({
     get data() {
      return data;
     },
     columns,
     getCoreRowModel: getCoreRowModel(),
     getPaginationRowModel: getPaginationRowModel(),
     getSortedRowModel: getSortedRowModel(),
     getFilteredRowModel: getFilteredRowModel(),
    onSortingChange: (updater) => {
      if (typeof updater === "function") {
        sorting = updater(sorting);
      } else {
        sorting = updater;
      }
    },
    onPaginationChange: (updater) => {
      if (typeof updater === "function") {
        pagination = updater(pagination);
      } else {
        pagination = updater;
      }
    },
    onColumnFiltersChange: (updater) => {
      if (typeof updater === "function") {
        columnFilters = updater(columnFilters);
      } else {
        columnFilters = updater;
      }
    },
    state: {
        get pagination() {
        return pagination;
      },
      get sorting() {
        return sorting;
      },
      get columnFilters() {
        return columnFilters;
      },
    },
    });


    // if you click a ctf you get redirected to the event room for that ctf
    const handleClick = (event: Event, row: { original: TData }) => {
        if (isChallenge(row.original)) {
            let ctf: Challenge = row.original;
            if (ctf.id.trim()) {
                goto(`/pages/event-room?code=${ctf.id}`);
            }
        }

        function isChallenge(data: any): data is Challenge {
            return (data as Challenge).id !== undefined;
        }

    }
   </script>


    <div>
        <div class="flex items-center py-4">
            <Input
              placeholder="Filter events by name..."
              value={(table.getColumn("name")?.getFilterValue() as string) ?? ""}
              onchange={(e: Event) => {
                if (e.currentTarget) {
                    table.getColumn("name")?.setFilterValue((e.currentTarget as HTMLInputElement).value);
                }
              }}
              oninput={(e: Event) => {
                if (e.currentTarget) {
                    table.getColumn("name")?.setFilterValue((e.currentTarget as HTMLInputElement).value);
                }
              }}
              class="max-w-sm"
            />
          </div>
        <div class="rounded-md border">
            <Table.Root>
            <Table.Header>
            {#each table.getHeaderGroups() as headerGroup (headerGroup.id)}
            <Table.Row>
                {#each headerGroup.headers as header (header.id)}
                <Table.Head>
                {#if !header.isPlaceholder}
                <FlexRender
                    content={header.column.columnDef.header}
                    context={header.getContext()}
                />
                {/if}
                </Table.Head>
                {/each}
            </Table.Row>
            {/each}
            </Table.Header>
            <Table.Body>
            {#each table.getRowModel().rows as row (row.id)}
            <Table.Row data-state={row.getIsSelected() && "selected"} onclick={(event)=>{handleClick(event, row)}}>
                {#each row.getVisibleCells() as cell (cell.id)}
                <Table.Cell>
                <FlexRender
                content={cell.column.columnDef.cell}
                context={cell.getContext()}
                />
                </Table.Cell>
                {/each}
            </Table.Row>
            {:else}
            <Table.Row>
                <Table.Cell colspan={columns.length} class="h-24 text-center">
                No results.
                </Table.Cell>
            </Table.Row>
            {/each}
            </Table.Body>
            </Table.Root>
        </div>
        <div class="flex items-center justify-end space-x-2 py-4">
            <Button
              variant="outline"
              size="sm"
              onclick={() => table.previousPage()}
              disabled={!table.getCanPreviousPage()}
            >
              Previous
            </Button>
            <Button
              variant="outline"
              size="sm"
              onclick={() => table.nextPage()}
              disabled={!table.getCanNextPage()}
            >
              Next
            </Button>
          </div>
    </div>
