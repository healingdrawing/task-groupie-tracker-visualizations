# GROUPIE-TRACKER FRONTEND

Next.js based frontend for [groupie-tracker](https://github.com/01-edu/public/tree/master/subjects/groupie-tracker) task

---

Authors: [@maximihajlov](https://github.com/maximihajlov), [@nattikim](https://github.com/nattikim)

Solved during studying in Gritlab coding school on Åland, November 2022

---

## Usage

You need Node.js installed to run frontend separately

### Run `npm install`

To get project dependencies

### Run `npm run dev`

To start development server

### Run `npm run build`

To build optimized production build

### Run `npm run start`

To start Next.js server
in [Incremental Static Regeneration](https://nextjs.org/docs/basic-features/data-fetching/incremental-static-regeneration)
mode

### or Run `npm run export`

To read from API only once. Static files will be saved to the directory `out`

[//]: # "TODO: add comment about data fetching and export mode config"

### Environment variables

You can customize API by changing following environment variables:

- `GROUPIE_BACKEND_LOCALHOST` - local host for server-side initial data fetching
- `NEXT_PUBLIC_GROUPIE_BACKEND_HOST` - API host for client-side data fetching
