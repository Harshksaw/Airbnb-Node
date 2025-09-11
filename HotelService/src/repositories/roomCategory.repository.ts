
import BaseRepository from "./base.repository";



/* This line of code is defining a class named `HotelRepository` that extends another class `BaseRepository` with a generic type `Hotel`. This means that `HotelRepository` inherits all the properties and methods from `BaseRepository` class and is specifically designed to work with objects of type `Hotel`. */
export class RoomCategoryRepository extends BaseRepository<RooomCategory> {



}