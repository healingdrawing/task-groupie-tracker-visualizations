import { Artist } from "../artist"
import useSWR from "swr"

export const getArtistsClient = (): Promise<Artist[]> =>
  //TODO: add export mode with ${process.env.NEXT_PUBLIC_GROUPIE_BACKEND_HOST}
  fetch(`/api/artists`)
    .then((res) => res.json())
    .then((data) => data as Artist[])

export const useArtists = () => {
  const { data, error } = useSWR("/artists/", getArtistsClient)
  return {
    artists: data,
    isLoading: !error && !data,
    isError: error as Error,
  }
}
