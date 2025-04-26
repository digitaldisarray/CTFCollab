import { S as escape_html, X as store_get, Y as unsubscribe_stores, t as pop, p as push, V as copy_payload, W as assign_payload, Z as bind_props, _ as spread_props, $ as ensure_array_like, a0 as spread_attributes, a4 as slot, a2 as once, a5 as sanitize_props, a9 as is_array, aa as get_prototype_of, ab as object_prototype, a6 as store_mutate } from './index-D3Djb8Yj.js';
import { B as Button, c as cn } from './button-dEtd2XS0.js';
import { I as Input } from './input-DTwr2lW9.js';
import { g as goto } from './client-CZszgVC4.js';
import { getCoreRowModel, getPaginationRowModel, getSortedRowModel, getFilteredRowModel } from '@tanstack/table-core';
import { c as createSvelteTable, T as Table, r as renderComponent, a as renderSnippet, b as Table_header, d as Table_body, R as Root$1, e as createRawSnippet, f as Table_row, g as Table_cell, i as Table_head, F as Flex_render, h as Trigger$1, D as Dropdown_menu_content, A as Arrow_up_down, I as Icon, G as Group, j as Dropdown_menu_separator, k as Dropdown_menu_item, E as Ellipsis, m as Dropdown_menu_group_heading } from './table-row-D24vYnD2.js';
import { p as page } from './stores-BfAkj-Gl.js';
import { w as writable } from './exports-wkiByGb4.js';
import { R as Root, T as Trigger, P as Popover_content, u as useId$1, b as box$1, m as mergeProps$1, C as Context, F as Form_field, a as Form_button, z as watch, E as ENTER, S as SPACE, s as srOnlyStylesString, c as Control, d as Form_description, e as Form_field_errors, B as getAriaChecked, D as getAriaRequired$1, o as getDataDisabled, v as Form_label } from './index3-VeECvYOq.js';
import { f as formSchema } from './5-nsEYY0ti.js';
import { b as superForm, c as zodClient } from './superValidate-C3ZLgWDn.js';
import './stringify-B14jp4IP.js';
import './events-DkGbpcOb.js';

const empty = [];
function snapshot(value, skip_warning = false) {
  return clone(value, /* @__PURE__ */ new Map(), "", empty);
}
function clone(value, cloned, path, paths, original = null) {
  if (typeof value === "object" && value !== null) {
    var unwrapped = cloned.get(value);
    if (unwrapped !== void 0) return unwrapped;
    if (value instanceof Map) return (
      /** @type {Snapshot<T>} */
      new Map(value)
    );
    if (value instanceof Set) return (
      /** @type {Snapshot<T>} */
      new Set(value)
    );
    if (is_array(value)) {
      var copy = (
        /** @type {Snapshot<any>} */
        Array(value.length)
      );
      cloned.set(value, copy);
      if (original !== null) {
        cloned.set(original, copy);
      }
      for (var i = 0; i < value.length; i += 1) {
        var element = value[i];
        if (i in value) {
          copy[i] = clone(element, cloned, path, paths);
        }
      }
      return copy;
    }
    if (get_prototype_of(value) === object_prototype) {
      copy = {};
      cloned.set(value, copy);
      if (original !== null) {
        cloned.set(original, copy);
      }
      for (var key in value) {
        copy[key] = clone(value[key], cloned, path, paths);
      }
      return copy;
    }
    if (value instanceof Date) {
      return (
        /** @type {Snapshot<T>} */
        structuredClone(value)
      );
    }
    if (typeof /** @type {T & { toJSON?: any } } */
    value.toJSON === "function") {
      return clone(
        /** @type {T & { toJSON(): any } } */
        value.toJSON(),
        cloned,
        path,
        paths,
        // Associate the instance with the toJSON clone
        value
      );
    }
  }
  if (value instanceof EventTarget) {
    return (
      /** @type {Snapshot<T>} */
      value
    );
  }
  try {
    return (
      /** @type {Snapshot<T>} */
      structuredClone(value)
    );
  } catch (e) {
    return (
      /** @type {Snapshot<T>} */
      value
    );
  }
}
const CHECKBOX_ROOT_ATTR = "data-checkbox-root";
class CheckboxRootState {
  opts;
  group;
  #trueName = once(() => {
    if (this.group && this.group.opts.name.current) {
      return this.group.opts.name.current;
    } else {
      return this.opts.name.current;
    }
  });
  get trueName() {
    return this.#trueName();
  }
  #trueRequired = once(() => {
    if (this.group && this.group.opts.required.current) {
      return true;
    }
    return this.opts.required.current;
  });
  get trueRequired() {
    return this.#trueRequired();
  }
  #trueDisabled = once(() => {
    if (this.group && this.group.opts.disabled.current) {
      return true;
    }
    return this.opts.disabled.current;
  });
  get trueDisabled() {
    return this.#trueDisabled();
  }
  constructor(opts, group = null) {
    this.opts = opts;
    this.group = group;
    this.onkeydown = this.onkeydown.bind(this);
    this.onclick = this.onclick.bind(this);
    watch.pre(
      [
        () => snapshot(this.group?.opts.value.current),
        () => this.opts.value.current
      ],
      ([groupValue, value]) => {
        if (!groupValue || !value) return;
        this.opts.checked.current = groupValue.includes(value);
      }
    );
    watch.pre(() => this.opts.checked.current, (checked) => {
      if (!this.group) return;
      if (checked) {
        this.group?.addValue(this.opts.value.current);
      } else {
        this.group?.removeValue(this.opts.value.current);
      }
    });
  }
  onkeydown(e) {
    if (this.opts.disabled.current) return;
    if (e.key === ENTER) e.preventDefault();
    if (e.key === SPACE) {
      e.preventDefault();
      this.#toggle();
    }
  }
  #toggle() {
    if (this.opts.indeterminate.current) {
      this.opts.indeterminate.current = false;
      this.opts.checked.current = true;
    } else {
      this.opts.checked.current = !this.opts.checked.current;
    }
  }
  onclick(_) {
    if (this.opts.disabled.current) return;
    this.#toggle();
  }
  #snippetProps = once(() => ({
    checked: this.opts.checked.current,
    indeterminate: this.opts.indeterminate.current
  }));
  get snippetProps() {
    return this.#snippetProps();
  }
  #props = once(() => ({
    id: this.opts.id.current,
    role: "checkbox",
    type: this.opts.type.current,
    disabled: this.trueDisabled,
    "aria-checked": getAriaChecked(this.opts.checked.current, this.opts.indeterminate.current),
    "aria-required": getAriaRequired$1(this.trueRequired),
    "data-disabled": getDataDisabled(this.trueDisabled),
    "data-state": getCheckboxDataState(this.opts.checked.current, this.opts.indeterminate.current),
    [CHECKBOX_ROOT_ATTR]: "",
    //
    onclick: this.onclick,
    onkeydown: this.onkeydown
  }));
  get props() {
    return this.#props();
  }
}
class CheckboxInputState {
  root;
  #trueChecked = once(() => {
    if (this.root.group) {
      if (this.root.opts.value.current !== void 0 && this.root.group.opts.value.current.includes(this.root.opts.value.current)) {
        return true;
      }
      return false;
    }
    return this.root.opts.checked.current;
  });
  get trueChecked() {
    return this.#trueChecked();
  }
  #shouldRender = once(() => Boolean(this.root.trueName));
  get shouldRender() {
    return this.#shouldRender();
  }
  constructor(root) {
    this.root = root;
  }
  #props = once(() => ({
    type: "checkbox",
    checked: this.root.opts.checked.current === true,
    disabled: this.root.trueDisabled,
    required: this.root.trueRequired,
    name: this.root.trueName,
    value: this.root.opts.value.current
  }));
  get props() {
    return this.#props();
  }
}
function getCheckboxDataState(checked, indeterminate) {
  if (indeterminate) return "indeterminate";
  return checked ? "checked" : "unchecked";
}
const CheckboxGroupContext = new Context("Checkbox.Group");
const CheckboxRootContext = new Context("Checkbox.Root");
function useCheckboxRoot(props, group) {
  return CheckboxRootContext.set(new CheckboxRootState(props, group));
}
function useCheckboxInput() {
  return new CheckboxInputState(CheckboxRootContext.get());
}
function Hidden_input($$payload, $$props) {
  push();
  let {
    value = void 0,
    $$slots,
    $$events,
    ...restProps
  } = $$props;
  const mergedProps = mergeProps$1(restProps, {
    "aria-hidden": "true",
    tabindex: -1,
    style: srOnlyStylesString
  });
  if (mergedProps.type === "checkbox") {
    $$payload.out += "<!--[-->";
    $$payload.out += `<input${spread_attributes({ ...mergedProps, value })}>`;
  } else {
    $$payload.out += "<!--[!-->";
    $$payload.out += `<input${spread_attributes({ value, ...mergedProps })}>`;
  }
  $$payload.out += `<!--]-->`;
  bind_props($$props, { value });
  pop();
}
function Checkbox_input($$payload, $$props) {
  push();
  const inputState = useCheckboxInput();
  if (inputState.shouldRender) {
    $$payload.out += "<!--[-->";
    Hidden_input($$payload, spread_props([inputState.props]));
  } else {
    $$payload.out += "<!--[!-->";
  }
  $$payload.out += `<!--]-->`;
  pop();
}
function Checkbox$1($$payload, $$props) {
  push();
  let {
    checked = false,
    ref = null,
    onCheckedChange,
    children,
    disabled = false,
    required = false,
    name = void 0,
    value = "on",
    id = useId$1(),
    indeterminate = false,
    onIndeterminateChange,
    child,
    type = "button",
    $$slots,
    $$events,
    ...restProps
  } = $$props;
  const group = CheckboxGroupContext.getOr(null);
  if (group && value) {
    if (group.opts.value.current.includes(value)) {
      checked = true;
    } else {
      checked = false;
    }
  }
  const rootState = useCheckboxRoot(
    {
      checked: box$1.with(() => checked, (v) => {
        checked = v;
        onCheckedChange?.(v);
      }),
      disabled: box$1.with(() => disabled ?? false),
      required: box$1.with(() => required),
      name: box$1.with(() => name),
      value: box$1.with(() => value),
      id: box$1.with(() => id),
      ref: box$1.with(() => ref, (v) => ref = v),
      indeterminate: box$1.with(() => indeterminate, (v) => {
        indeterminate = v;
        onIndeterminateChange?.(v);
      }),
      type: box$1.with(() => type)
    },
    group
  );
  const mergedProps = mergeProps$1({ ...restProps }, rootState.props);
  if (child) {
    $$payload.out += "<!--[-->";
    child($$payload, { props: mergedProps, ...rootState.snippetProps });
    $$payload.out += `<!---->`;
  } else {
    $$payload.out += "<!--[!-->";
    $$payload.out += `<button${spread_attributes({ ...mergedProps })}>`;
    children?.($$payload, rootState.snippetProps);
    $$payload.out += `<!----></button>`;
  }
  $$payload.out += `<!--]--> `;
  Checkbox_input($$payload);
  $$payload.out += `<!---->`;
  bind_props($$props, { checked, ref, indeterminate });
  pop();
}
function Check($$payload, $$props) {
  const $$sanitized_props = sanitize_props($$props);
  const iconNode = [["path", { "d": "M20 6 9 17l-5-5" }]];
  Icon($$payload, spread_props([
    { name: "check" },
    $$sanitized_props,
    {
      iconNode,
      children: ($$payload2) => {
        $$payload2.out += `<!---->`;
        slot($$payload2, $$props, "default", {});
        $$payload2.out += `<!---->`;
      },
      $$slots: { default: true }
    }
  ]));
}
function Minus($$payload, $$props) {
  const $$sanitized_props = sanitize_props($$props);
  const iconNode = [["path", { "d": "M5 12h14" }]];
  Icon($$payload, spread_props([
    { name: "minus" },
    $$sanitized_props,
    {
      iconNode,
      children: ($$payload2) => {
        $$payload2.out += `<!---->`;
        slot($$payload2, $$props, "default", {});
        $$payload2.out += `<!---->`;
      },
      $$slots: { default: true }
    }
  ]));
}
function Checkbox($$payload, $$props) {
  push();
  let {
    ref = null,
    checked = false,
    indeterminate = false,
    class: className,
    $$slots,
    $$events,
    ...restProps
  } = $$props;
  let $$settled = true;
  let $$inner_payload;
  function $$render_inner($$payload2) {
    $$payload2.out += `<!---->`;
    {
      let children = function($$payload3, { checked: checked2, indeterminate: indeterminate2 }) {
        $$payload3.out += `<div class="flex size-4 items-center justify-center text-current">`;
        if (indeterminate2) {
          $$payload3.out += "<!--[-->";
          Minus($$payload3, { class: "size-3.5" });
        } else {
          $$payload3.out += "<!--[!-->";
          Check($$payload3, {
            class: cn("size-3.5", !checked2 && "text-transparent")
          });
        }
        $$payload3.out += `<!--]--></div>`;
      };
      Checkbox$1($$payload2, spread_props([
        {
          class: cn("border-primary ring-offset-background focus-visible:ring-ring data-[state=checked]:bg-primary data-[state=checked]:text-primary-foreground peer box-content size-4 shrink-0 rounded-sm border focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 data-[disabled=true]:cursor-not-allowed data-[disabled=true]:opacity-50", className)
        },
        restProps,
        {
          get ref() {
            return ref;
          },
          set ref($$value) {
            ref = $$value;
            $$settled = false;
          },
          get checked() {
            return checked;
          },
          set checked($$value) {
            checked = $$value;
            $$settled = false;
          },
          get indeterminate() {
            return indeterminate;
          },
          set indeterminate($$value) {
            indeterminate = $$value;
            $$settled = false;
          },
          children,
          $$slots: { default: true }
        }
      ]));
    }
    $$payload2.out += `<!---->`;
  }
  do {
    $$settled = true;
    $$inner_payload = copy_payload($$payload);
    $$render_inner($$inner_payload);
  } while (!$$settled);
  assign_payload($$payload, $$inner_payload);
  bind_props($$props, { ref, checked, indeterminate });
  pop();
}
function Data_table($$payload, $$props) {
  push();
  var $$store_subs;
  let { data, columns: columns2 } = $$props;
  let pagination = { pageIndex: 0, pageSize: 10 };
  let sorting = [];
  let columnFilters = [];
  const table = createSvelteTable({
    get data() {
      return data;
    },
    columns: columns2,
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
      }
    }
  });
  const handleClick = (event, row) => {
    console.log("Clicked row:", row.original);
    if (isChallenge(row.original)) {
      let challenge = row.original;
      console.log("Challenge ID:", challenge.id);
      const ctfCode = new URLSearchParams(store_get($$store_subs ??= {}, "$page", page).url.search).get("code");
      console.log("CTF Code:", ctfCode);
      if (!ctfCode) {
        console.error("CTF code not found in URL");
        return;
      }
      if (challenge.id.trim()) {
        const url = `/pages/event-room/challenge?code=${encodeURIComponent(ctfCode)}&challenge=${challenge.id}`;
        console.log("Navigating to:", url);
        goto();
      }
    }
    function isChallenge(data2) {
      return data2.id !== void 0;
    }
  };
  $$payload.out += `<div><div class="flex items-center py-4">`;
  Input($$payload, {
    placeholder: "Filter challenge by name...",
    value: table.getColumn("name")?.getFilterValue() ?? "",
    onchange: (e) => {
      table.getColumn("name")?.setFilterValue(e.currentTarget.value);
    },
    oninput: (e) => {
      table.getColumn("name")?.setFilterValue(e.currentTarget.value);
    },
    class: "max-w-sm"
  });
  $$payload.out += `<!----></div> <div class="rounded-md border"><!---->`;
  Table($$payload, {
    children: ($$payload2) => {
      $$payload2.out += `<!---->`;
      Table_header($$payload2, {
        children: ($$payload3) => {
          const each_array = ensure_array_like(table.getHeaderGroups());
          $$payload3.out += `<!--[-->`;
          for (let $$index_1 = 0, $$length = each_array.length; $$index_1 < $$length; $$index_1++) {
            let headerGroup = each_array[$$index_1];
            $$payload3.out += `<!---->`;
            Table_row($$payload3, {
              children: ($$payload4) => {
                const each_array_1 = ensure_array_like(headerGroup.headers);
                $$payload4.out += `<!--[-->`;
                for (let $$index = 0, $$length2 = each_array_1.length; $$index < $$length2; $$index++) {
                  let header = each_array_1[$$index];
                  $$payload4.out += `<!---->`;
                  Table_head($$payload4, {
                    children: ($$payload5) => {
                      if (!header.isPlaceholder) {
                        $$payload5.out += "<!--[-->";
                        Flex_render($$payload5, {
                          content: header.column.columnDef.header,
                          context: header.getContext()
                        });
                      } else {
                        $$payload5.out += "<!--[!-->";
                      }
                      $$payload5.out += `<!--]-->`;
                    },
                    $$slots: { default: true }
                  });
                  $$payload4.out += `<!---->`;
                }
                $$payload4.out += `<!--]-->`;
              },
              $$slots: { default: true }
            });
            $$payload3.out += `<!---->`;
          }
          $$payload3.out += `<!--]-->`;
        },
        $$slots: { default: true }
      });
      $$payload2.out += `<!----> <!---->`;
      Table_body($$payload2, {
        children: ($$payload3) => {
          const each_array_2 = ensure_array_like(table.getRowModel().rows);
          if (each_array_2.length !== 0) {
            $$payload3.out += "<!--[-->";
            for (let $$index_3 = 0, $$length = each_array_2.length; $$index_3 < $$length; $$index_3++) {
              let row = each_array_2[$$index_3];
              $$payload3.out += `<!---->`;
              Table_row($$payload3, {
                "data-state": row.getIsSelected() && "selected",
                onclick: (event) => handleClick(event, row),
                children: ($$payload4) => {
                  const each_array_3 = ensure_array_like(row.getVisibleCells());
                  $$payload4.out += `<!--[-->`;
                  for (let $$index_2 = 0, $$length2 = each_array_3.length; $$index_2 < $$length2; $$index_2++) {
                    let cell = each_array_3[$$index_2];
                    $$payload4.out += `<!---->`;
                    Table_cell($$payload4, {
                      children: ($$payload5) => {
                        Flex_render($$payload5, {
                          content: cell.column.columnDef.cell,
                          context: cell.getContext()
                        });
                      },
                      $$slots: { default: true }
                    });
                    $$payload4.out += `<!---->`;
                  }
                  $$payload4.out += `<!--]-->`;
                },
                $$slots: { default: true }
              });
              $$payload3.out += `<!---->`;
            }
          } else {
            $$payload3.out += "<!--[!-->";
            $$payload3.out += `<!---->`;
            Table_row($$payload3, {
              children: ($$payload4) => {
                $$payload4.out += `<!---->`;
                Table_cell($$payload4, {
                  colspan: columns2.length,
                  class: "h-24 text-center",
                  children: ($$payload5) => {
                    $$payload5.out += `<!---->No results.`;
                  },
                  $$slots: { default: true }
                });
                $$payload4.out += `<!---->`;
              },
              $$slots: { default: true }
            });
            $$payload3.out += `<!---->`;
          }
          $$payload3.out += `<!--]-->`;
        },
        $$slots: { default: true }
      });
      $$payload2.out += `<!---->`;
    },
    $$slots: { default: true }
  });
  $$payload.out += `<!----></div> <div class="flex items-center justify-end space-x-2 py-4">`;
  Button($$payload, {
    variant: "outline",
    size: "sm",
    onclick: () => table.previousPage(),
    disabled: !table.getCanPreviousPage(),
    children: ($$payload2) => {
      $$payload2.out += `<!---->Previous`;
    },
    $$slots: { default: true }
  });
  $$payload.out += `<!----> `;
  Button($$payload, {
    variant: "outline",
    size: "sm",
    onclick: () => table.nextPage(),
    disabled: !table.getCanNextPage(),
    children: ($$payload2) => {
      $$payload2.out += `<!---->Next`;
    },
    $$slots: { default: true }
  });
  $$payload.out += `<!----></div></div>`;
  if ($$store_subs) unsubscribe_stores($$store_subs);
  pop();
}
function Data_table_actions($$payload, $$props) {
  push();
  var $$store_subs;
  const ctfPhrase = new URLSearchParams(store_get($$store_subs ??= {}, "$page", page).url.search).get("code");
  let { id, description } = $$props;
  const deleteChal = async () => {
    const confirmed = confirm("Are you sure you want to delete this Challenge? This action is irreversible.");
    if (!confirmed) return;
    try {
      const response = await fetch(`http://localhost:1337/ctfs/${ctfPhrase}/challenges/${id}`, { method: "DELETE", credentials: "include" });
      if (response.ok) {
        challenges.update((current) => current.filter((ch) => ch.id !== id));
        alert("Challenge deleted successfully");
      } else {
        console.error("Delete request failed:", response);
        alert("Failed to delete the Challenge.");
      }
    } catch (error) {
      console.error("Error deleting Challenge:", error);
      alert("An error occurred while deleting the Challenge.");
    }
  };
  $$payload.out += `<!---->`;
  Root$1($$payload, {
    children: ($$payload2) => {
      $$payload2.out += `<!---->`;
      {
        let child = function($$payload3, { props }) {
          Button($$payload3, spread_props([
            props,
            {
              variant: "ghost",
              size: "icon",
              class: "relative size-8 p-0",
              children: ($$payload4) => {
                $$payload4.out += `<span class="sr-only">Open menu</span> `;
                Ellipsis($$payload4, {});
                $$payload4.out += `<!---->`;
              },
              $$slots: { default: true }
            }
          ]));
        };
        Trigger$1($$payload2, { child, $$slots: { child: true } });
      }
      $$payload2.out += `<!----> <!---->`;
      Dropdown_menu_content($$payload2, {
        children: ($$payload3) => {
          $$payload3.out += `<!---->`;
          Group($$payload3, {
            children: ($$payload4) => {
              $$payload4.out += `<!---->`;
              Dropdown_menu_group_heading($$payload4, {
                children: ($$payload5) => {
                  $$payload5.out += `<!---->Actions`;
                },
                $$slots: { default: true }
              });
              $$payload4.out += `<!----> <!---->`;
              Dropdown_menu_item($$payload4, {
                onclick: () => navigator.clipboard.writeText(id),
                children: ($$payload5) => {
                  $$payload5.out += `<!---->Copy Challenge ID`;
                },
                $$slots: { default: true }
              });
              $$payload4.out += `<!---->`;
            },
            $$slots: { default: true }
          });
          $$payload3.out += `<!----> <!---->`;
          Dropdown_menu_separator($$payload3, {});
          $$payload3.out += `<!----> <!---->`;
          Dropdown_menu_item($$payload3, {
            children: ($$payload4) => {
              $$payload4.out += `<!---->View Active Members`;
            },
            $$slots: { default: true }
          });
          $$payload3.out += `<!----> <!---->`;
          Dropdown_menu_item($$payload3, {
            onclick: () => navigator.clipboard.writeText(description),
            children: ($$payload4) => {
              $$payload4.out += `<!---->View Challenge details`;
            },
            $$slots: { default: true }
          });
          $$payload3.out += `<!----> <!---->`;
          Dropdown_menu_item($$payload3, {
            onclick: deleteChal,
            children: ($$payload4) => {
              $$payload4.out += `<!---->Delete Challenge`;
            },
            $$slots: { default: true }
          });
          $$payload3.out += `<!---->`;
        },
        $$slots: { default: true }
      });
      $$payload2.out += `<!---->`;
    },
    $$slots: { default: true }
  });
  $$payload.out += `<!---->`;
  if ($$store_subs) unsubscribe_stores($$store_subs);
  pop();
}
function Data_name_button($$payload, $$props) {
  let {
    variant = "ghost",
    $$slots,
    $$events,
    ...restProps
  } = $$props;
  Button($$payload, spread_props([
    { variant },
    restProps,
    {
      children: ($$payload2) => {
        $$payload2.out += `<!---->Name `;
        Arrow_up_down($$payload2, { class: "ml-2" });
        $$payload2.out += `<!---->`;
      },
      $$slots: { default: true }
    }
  ]));
}
const challenges = writable([]);
const currentCTF = writable();
const columns = [
  {
    accessorKey: "status",
    header: "Status",
    cell: ({ row }) => renderComponent(Checkbox, {
      checked: row.getIsSelected(),
      onCheckedChange: (value) => row.toggleSelected(!!value),
      "aria-label": "Select row"
    }),
    enableSorting: false,
    enableHiding: false
  },
  {
    accessorKey: "name",
    header: ({ column }) => renderComponent(Data_name_button, {
      onclick: () => column.toggleSorting(column.getIsSorted() === "asc")
    })
  },
  {
    accessorKey: "active_members",
    header: () => {
      const active_membersHeaderSnippet = createRawSnippet(() => ({
        render: () => `<div class="text-right">Active Members</div>`
      }));
      return renderSnippet(active_membersHeaderSnippet, "");
    },
    cell: ({ row }) => {
      const formatter = new Intl.NumberFormat("en-US", {
        style: "decimal"
      });
      const active_membersCellSnippet = createRawSnippet((getAmount) => {
        const active_members = getAmount();
        return {
          render: () => `<div class="text-right font-medium">${active_members}</div>`
        };
      });
      return renderSnippet(
        active_membersCellSnippet,
        formatter.format(parseFloat(row.getValue("active_members")))
      );
    }
  },
  {
    id: "actions",
    cell: ({ row }) => {
      return renderComponent(Data_table_actions, {
        id: row.original.id,
        description: row.original.description
      });
    }
  }
];
function Settings_form($$payload, $$props) {
  push();
  var $$store_subs;
  let { data } = $$props;
  const form = superForm(data.form, { validators: zodClient(formSchema) });
  const { form: formData, enhance } = form;
  let $$settled = true;
  let $$inner_payload;
  function $$render_inner($$payload2) {
    $$payload2.out += `<form method="POST"><!---->`;
    Form_field($$payload2, {
      form,
      name: "challengeName",
      children: ($$payload3) => {
        $$payload3.out += `<!---->`;
        {
          let children = function($$payload4, { props }) {
            $$payload4.out += `<!---->`;
            Form_label($$payload4, {
              children: ($$payload5) => {
                $$payload5.out += `<!---->Challenge Name`;
              },
              $$slots: { default: true }
            });
            $$payload4.out += `<!----> `;
            Input($$payload4, spread_props([
              props,
              {
                get value() {
                  return store_get($$store_subs ??= {}, "$formData", formData).challengeName;
                },
                set value($$value) {
                  store_mutate($$store_subs ??= {}, "$formData", formData, store_get($$store_subs ??= {}, "$formData", formData).challengeName = $$value);
                  $$settled = false;
                }
              }
            ]));
            $$payload4.out += `<!---->`;
          };
          Control($$payload3, { children, $$slots: { default: true } });
        }
        $$payload3.out += `<!----> <!---->`;
        Form_description($$payload3, {
          children: ($$payload4) => {
            $$payload4.out += `<!---->This sets the Challenge display name.`;
          },
          $$slots: { default: true }
        });
        $$payload3.out += `<!----> <!---->`;
        Form_field_errors($$payload3, {});
        $$payload3.out += `<!---->`;
      },
      $$slots: { default: true }
    });
    $$payload2.out += `<!----> <!---->`;
    Form_field($$payload2, {
      form,
      name: "challengeDetails",
      children: ($$payload3) => {
        $$payload3.out += `<!---->`;
        {
          let children = function($$payload4, { props }) {
            $$payload4.out += `<!---->`;
            Form_label($$payload4, {
              children: ($$payload5) => {
                $$payload5.out += `<!---->(Optional) Challenge Details`;
              },
              $$slots: { default: true }
            });
            $$payload4.out += `<!----> `;
            Input($$payload4, spread_props([
              props,
              {
                get value() {
                  return store_get($$store_subs ??= {}, "$formData", formData).challengeDetails;
                },
                set value($$value) {
                  store_mutate($$store_subs ??= {}, "$formData", formData, store_get($$store_subs ??= {}, "$formData", formData).challengeDetails = $$value);
                  $$settled = false;
                }
              }
            ]));
            $$payload4.out += `<!---->`;
          };
          Control($$payload3, { children, $$slots: { default: true } });
        }
        $$payload3.out += `<!----> <!---->`;
        Form_description($$payload3, {
          children: ($$payload4) => {
            $$payload4.out += `<!---->Describe the details of this challenge`;
          },
          $$slots: { default: true }
        });
        $$payload3.out += `<!----> <!---->`;
        Form_field_errors($$payload3, {});
        $$payload3.out += `<!---->`;
      },
      $$slots: { default: true }
    });
    $$payload2.out += `<!----> <!---->`;
    Form_field($$payload2, {
      form,
      name: "challengeFlag",
      children: ($$payload3) => {
        $$payload3.out += `<!---->`;
        {
          let children = function($$payload4, { props }) {
            $$payload4.out += `<!---->`;
            Form_label($$payload4, {
              children: ($$payload5) => {
                $$payload5.out += `<!---->(Optional) Challenge Flag`;
              },
              $$slots: { default: true }
            });
            $$payload4.out += `<!----> `;
            Input($$payload4, spread_props([
              props,
              {
                get value() {
                  return store_get($$store_subs ??= {}, "$formData", formData).challengeFlag;
                },
                set value($$value) {
                  store_mutate($$store_subs ??= {}, "$formData", formData, store_get($$store_subs ??= {}, "$formData", formData).challengeFlag = $$value);
                  $$settled = false;
                }
              }
            ]));
            $$payload4.out += `<!---->`;
          };
          Control($$payload3, { children, $$slots: { default: true } });
        }
        $$payload3.out += `<!----> <!---->`;
        Form_description($$payload3, {
          children: ($$payload4) => {
            $$payload4.out += `<!---->The flag for this challenge`;
          },
          $$slots: { default: true }
        });
        $$payload3.out += `<!----> <!---->`;
        Form_field_errors($$payload3, {});
        $$payload3.out += `<!---->`;
      },
      $$slots: { default: true }
    });
    $$payload2.out += `<!----> <!---->`;
    Form_button($$payload2, {
      children: ($$payload3) => {
        $$payload3.out += `<!---->Submit`;
      },
      $$slots: { default: true }
    });
    $$payload2.out += `<!----></form>`;
  }
  do {
    $$settled = true;
    $$inner_payload = copy_payload($$payload);
    $$render_inner($$inner_payload);
  } while (!$$settled);
  assign_payload($$payload, $$inner_payload);
  if ($$store_subs) unsubscribe_stores($$store_subs);
  pop();
}
function _page($$payload, $$props) {
  push();
  var $$store_subs;
  let { data: pageData } = $$props;
  $$payload.out += `<header class="logo-header svelte-1tu79z4"><div class="absolute left-4 top-4 md:left-8 md:top-7"><a href="/" class="logo svelte-1tu79z4"><span class="ctf svelte-1tu79z4">CTF</span> <span class="collab svelte-1tu79z4">Collab</span></a></div></header> <div class="page-container svelte-1tu79z4"><div class="container svelte-1tu79z4"><div class="horizontal-container svelte-1tu79z4"><header class="svelte-1tu79z4"><h1 class="chal-name svelte-1tu79z4">${escape_html(store_get($$store_subs ??= {}, "$currentCTF", currentCTF)?.ctf_name)}</h1> <p class="svelte-1tu79z4">Manage your Challenge rooms.</p></header> <div class="button-container svelte-1tu79z4"><!---->`;
  Root($$payload, {
    children: ($$payload2) => {
      $$payload2.out += `<!---->`;
      Trigger($$payload2, {
        children: ($$payload3) => {
          Button($$payload3, {
            class: "new-ctf-button",
            children: ($$payload4) => {
              $$payload4.out += `<!---->Add Challenge`;
            },
            $$slots: { default: true }
          });
        },
        $$slots: { default: true }
      });
      $$payload2.out += `<!----> <!---->`;
      Popover_content($$payload2, {
        children: ($$payload3) => {
          Settings_form($$payload3, { data: pageData });
        },
        $$slots: { default: true }
      });
      $$payload2.out += `<!---->`;
    },
    $$slots: { default: true }
  });
  $$payload.out += `<!----></div></div> `;
  Data_table($$payload, {
    data: store_get($$store_subs ??= {}, "$challenges", challenges),
    columns
  });
  $$payload.out += `<!----></div> <div class="container svelte-1tu79z4"><header class="svelte-1tu79z4"><h1 class="chal-name svelte-1tu79z4">Team Members</h1> <p class="svelte-1tu79z4">Invite team members to collaborate on challenges.</p></header></div></div>`;
  if ($$store_subs) unsubscribe_stores($$store_subs);
  pop();
}

export { _page as default };
//# sourceMappingURL=_page.svelte-D7lwqf4H.js.map
