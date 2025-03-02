import { z } from "zod";
 
export const formSchema = z.object({
  ctfName: z.string().min(2).max(50),
  dob: z
      .string()
      .refine((v) => v, { message: "A scheduled event date is required" }),
  ctfCode: z.string().min(2).max(50),


});
 
export type FormSchema = typeof formSchema;