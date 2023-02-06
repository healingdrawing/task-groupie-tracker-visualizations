export type Artist = {
  id: number
  image: string
  name: string
  members: string[]
  creationDate: string
  firstAlbum: string
  cities: City[]
}

export type City = {
  name: string
  dates: string[]
}
