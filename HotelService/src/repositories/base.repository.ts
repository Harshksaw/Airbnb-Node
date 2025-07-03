import { Model, CreationAttributes, ModelStatic, WhereOptions } from 'sequelize';

class BaseRepository<T extends Model> {
        /**
        * The model that this repository is responsible for.
        * It should be a Sequelize model class.
        */
    protected model: ModelStatic<T>;

    constructor(model: ModelStatic<T>) {
        this.model = model;
    }
    async findById(id: number): Promise<T | null> {
        const record  = await this.model.findByPk(id);
        if (!record) {
            throw new Error(`Record with id ${id} not found`);
        }
        return record;
    }

     async findAll(): Promise<T[]> {
        const records = await this.model.findAll({
            where:{
                whereOptions: {

                    deletedAt: null // Assuming soft delete is implemented with a deletedAt field
                }
            }
        });
        return records;

     }
     async delete(whereOptions: WhereOptions<T>): Promise<void> {
        const record =await this.model.destroy({
            where:{
                ...whereOptions
            }

        })
        if(!record) {
            throw new Error(`Record with id ${JSON.stringify(whereOptions)} not found`);
        }
        return 

     }

     async create(data: CreationAttributes<T>): Promise<T> {



     }

     async update(id: number, data: Partial<T>): Promise<T | null> {}
}


export default BaseRepository;

