module Api
  class TodosController < ApplicationController
      def index
          render json: Todo.all
      end
      def create
          print "Json is: "
          print params
          @u = User.find_by_id(params['user_id'])
          if @u == nil
              render json: {message: "User not found"}
          end
          @t = Todo.new
          @t.task = params['task']
          @t.user = @u
          if @t.save
              render json: @t
          else
              render nothing: true, status: :bad_request
          end
      end
      def show
          print params['id']
          @t = Todo.find_by_id(params['id'])
          render json: @t
      end
      def new
          create()
      end
      def destroy
          @t = Todo.find_by_id(params['id'])
          @t.destroy()
          render json: {message: "todo deleted"}
      end
      def edit
          @t = Todo.find_by_id(params['id'])
          @u = User.find_by_id(params['user_id'])
          @t.task = params['task']
          @t.user = @u
          if @t.save
              render json: @t
          else
              render nothing: true, status: :bad_request
          end
      end
      def update
          edit()
      end

  end
end
