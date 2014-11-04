var models = require('../models');
var express = require('express');
var logger = require('morgan');
var router = express.Router();

router.use(logger());
/* GET home page. */
router.get('/', function(req, res) {
  res.render('index', { title: 'Express' });
});

router.get('/version', function(req, res) {
    res.json({v: 1.0, type: 'node.js'})
});

router.route('/users')
    .get(function(req, res) {
        models.User.findAll().success(function(users) {
            if(users) {
                res.json(users);
            } else {
                res.send(401, "User not found");
            }
        }).error(function(error){
            res.send(500, error);
        });
    })
    .post(function(req, res) {
        var user = models.User.build();
        user.username = req.body.username;
        user.save().success(function(success) {
            res.json(user);
        }).error(function(error) {
            res.send(500, error);
        });
    });

router.route('/users/:user_id')
    .get(function(req, res) {
        models.User.find({where: {id: req.params.user_id}})
            .success(function(users) {
                if(users){
                    res.json(users);
                } else {
                    res.send(401, "User not found");
                }
            }).error(function(error) {
                res.send(500, error);
            })
    })
    .put(function(req, res) {
        models.User.update({username: req.body.username}, {id: req.params.user_id})
            .success(function(success) {
                if(success) {
                    res.json({message: 'User updated!'});
                } else {
                    res.send(401, "User not found!");
                }
            }).error(function(error) {
                res.send(500, error);
            });
    })
    .delete(function(req, res){
        models.User.destroy({where: {id: req.params.user_id}}).success(function(success){
            if(success){
                res.json({message: "User deleted"});
            } else {
                res.send(404, "User not found");
            }
        }).error(function(error){
            res.send(500, error);
        })
    });
router.route('/users/:user_id/tasks')
    .get(function(req, res) {
        models.Task.findAll({where: {UserId: req.params.user_id}})
            .success(function(tasks) {
                if(tasks) {
                    res.json(tasks);
                } else {
                    res.send(404, "Tasks not found");
                }
            })
            .error(function(error){
                res.send(500, error);
            });
    });
router.route('/tasks')
    .get(function(req, res) {
        models.Task.findAll().success(function(tasks){
            if(tasks) {
                res.json(tasks);
            } else {
                res.send(401, "No tasks found");
            }
        }).error(function(error) {
            res.send(500, error);
        });
    })
    .post(function(req, res) {
        var task = models.Task.build();
        task.title = req.body.title;
        task.UserId = req.body.user_id;
        task.save().success(function(success){
            if(success){
                res.json({message: "Task saved!"});
            } else {
                res.send(500, "Task not saved!");
            }
        })
        .error(function(error) {
            res.send(500, error);
        });
    });

router.route('/tasks/:task_id')
    .get(function(req, res) {
        models.Task.find({where: {id: req.params.task_id}}).success(function(success){
            if(success){
                res.json(success);
            } else {
                res.send(401, "Task not found");
            }
        }).error(function(error){
            res.send(500, error);
        });
    })
    .put(function(req, res) {
        console.log(req.body.UserId);
        models.User.findOne(
                {
                    where: {
                            id: req.body.UserId
                            }
                })
                .success(function(user){
                    if(user){
                        models
                            .Task
                            .update(
                                {
                                    title: req.body.title,
                                    UserId: req.body.UserId
                                },
                                {
                                    where: {
                                            id: req.params.task_id
                                            }
                                })
                                .success(function(success){
                                    if(success){
                                        res.json({message: "Task updated!"});
                                    } else {
                                        res.send(404, "Task not found");
                                    }
                                })
                                .error(function(error) {
                                    res.send(500, error);
                                });
                    } else {
                        res.send(404, "User not found");
                    }
                }).error(function(error) {
                    res.send(500, error);
                });
    })
    .delete(function(req, res) {
        models.Task.destroy({where: {id: req.params.task_id}}).success(function(success){
            if(success){
                res.json({message: "Task deleted"});
            } else {
                res.send(404, "Task not found");
            }
        }).error(function(error){
            res.send(500, error);
        });
    });
module.exports = router;
