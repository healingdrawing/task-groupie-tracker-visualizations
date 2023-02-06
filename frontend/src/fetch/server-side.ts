import { Artist } from "../artist"

export const getArtistsLocal = (): Promise<Artist[]> =>
  fetch(`http://${process.env.GROUPIE_BACKEND_LOCALHOST}/api/artists/`)
    .then((res) => res.json())
    .then((data) => data as Artist[])

export const getArtistById = (id: number): Promise<Artist> =>
  fetch(`http://${process.env.GROUPIE_BACKEND_LOCALHOST}/api/artists/${id}`)
    .then((res) => res.json())
    .then((data) => data as Artist)
export const getArtistByName = (name: string): Promise<Artist | undefined> =>
  fetch(`http://${process.env.GROUPIE_BACKEND_LOCALHOST}/api/artists/?name=${name}`)
    .then((res) => {
      if (!res.ok) {
        return undefined
      }
      return res.json()
    })
    .then((data) => data as Artist)
