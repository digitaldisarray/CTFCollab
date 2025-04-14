import { p as push, a0 as spread_attributes, a1 as clsx, Z as bind_props, t as pop } from './index-D3Djb8Yj.js';
import { c as cn } from './button-dEtd2XS0.js';

function Input($$payload, $$props) {
  push();
  let {
    ref = null,
    value = void 0,
    class: className,
    $$slots,
    $$events,
    ...restProps
  } = $$props;
  $$payload.out += `<input${spread_attributes({
    class: clsx(cn("border-input bg-background ring-offset-background placeholder:text-muted-foreground focus-visible:ring-ring flex h-10 w-full rounded-md border px-3 py-2 text-base file:border-0 file:bg-transparent file:text-sm file:font-medium focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 md:text-sm", className)),
    value,
    ...restProps
  })}>`;
  bind_props($$props, { ref, value });
  pop();
}

export { Input as I };
//# sourceMappingURL=input-DTwr2lW9.js.map
