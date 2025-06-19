import fs from 'fs/promises';
import path from 'path';
import Handlebars from 'handlebars';


export async function renderMailTemplate(templateId: string, params: Record<string, any>): Promise<string> {
    const templatePath = path.join(__dirname, 'mailer', `${templateId}.hbs`);

    try{
        const content = await fs.readFile(templatePath, 'utf8');

        const finalContent = Handlebars.compile(content);

        return finalContent(params);
    }
    catch(err) {
        console.error(`Error accessing template ${templateId}:`, err);
        throw new Error(`Template not found: ${templateId}`);
    }

   
}