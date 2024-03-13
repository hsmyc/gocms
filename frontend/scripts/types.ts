export type SingleType = {
  name: string;
  fields: Field;
};

export type Field = {
  name: string;
  type?: string;
}[];
