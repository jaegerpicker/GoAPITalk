odule Api
  class TodosController < Api::BaseController

    private

      def todo_params
        params.require(:todo).permit(:task)
      end

      def query_params
        params.permit(:task)
      end

  end
end
