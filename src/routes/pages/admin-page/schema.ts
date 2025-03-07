import { z } from "zod";
 
export const formSchema = z.object({
  ctfName: z.string().min(2).max(50),
  start_date: z
      .string()
      .refine((v) => v, { message: "A scheduled start date is required" }),
  end_date: z
      .string()
      .refine((v) => v, { message: "A scheduled end date is required" }),
  ctfDescription: z.string().min(0).max(50),
  
});
 
export type FormSchema = typeof formSchema;